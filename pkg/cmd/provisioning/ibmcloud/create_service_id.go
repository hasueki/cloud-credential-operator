package ibmcloud

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"

	credreqv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning"
	"github.com/openshift/cloud-credential-operator/pkg/ibmcloud"
)

// APIKeyEnvVars is a list of environment variable names containing an IBM Cloud API key
var APIKeyEnvVars = []string{"IC_API_KEY", "IBMCLOUD_API_KEY", "BM_API_KEY", "BLUEMIX_API_KEY"}

var (
	// CreateOpts captures the options that affect creation of the generated
	// objects.
	CreateOpts = options{
		TargetDir: "",
	}

	// APIKeyName is the name for the autogenerated API Key for the Service ID
	APIKeyName = "ccoctl-generated-key"
)

// getEnv reads the content from first found environment variable from the envs list, returns empty string if none found.
func getEnv(envs []string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}

// NewCreateServiceIDCmd provides the "create-service-id" subcommand
func NewCreateServiceIDCmd() *cobra.Command {
	createServiceIDCmd := &cobra.Command{
		Use:              "create-service-id",
		Short:            "Create Service ID",
		RunE:             createServiceIDCmd,
		PersistentPreRun: initEnvForCreateServiceIDCmd,
	}

	createServiceIDCmd.PersistentFlags().StringVar(&CreateOpts.Name, "name", "", "User-defined name for all created IBM Cloud resources (can be separate from the cluster's infra-id)")
	createServiceIDCmd.MarkPersistentFlagRequired("name")
	createServiceIDCmd.PersistentFlags().StringVar(&CreateOpts.CredRequestDir, "credentials-requests-dir", "", "Directory containing files of CredentialsRequests to create IAM Roles for (can be created by running 'oc adm release extract --credentials-requests --cloud=aws' against an OpenShift release image)")
	createServiceIDCmd.MarkPersistentFlagRequired("credentials-requests-dir")
	createServiceIDCmd.PersistentFlags().StringVar(&CreateOpts.ResourceGroupName, "resource-group-name", "", "Name of the resource group used for scoping the access policies")
	createServiceIDCmd.PersistentFlags().StringVar(&CreateOpts.TargetDir, "output-dir", "", "Directory to place generated files (defaults to current directory)")

	return createServiceIDCmd
}

func createServiceIDCmd(cmd *cobra.Command, args []string) error {
	apiKey := getEnv(APIKeyEnvVars)
	if apiKey == "" {
		return fmt.Errorf("%s environment variable not set", APIKeyEnvVars)
	}

	params := &ibmcloud.ClientParams{
		InfraName: CreateOpts.Name,
	}

	ibmclient, err := ibmcloud.NewClient(apiKey, params)
	if err != nil {
		return err
	}

	apiKeyDetailsOptions := ibmclient.NewGetAPIKeysDetailsOptions()
	apiKeyDetailsOptions.SetIamAPIKey(apiKey)
	apiKeyDetails, _, err := ibmclient.GetAPIKeysDetails(apiKeyDetailsOptions)
	if err != nil {
		return errors.Wrap(err, "Failed to get Details for the given APIKey")
	}

	err = createServiceIDs(ibmclient, apiKeyDetails.AccountID, CreateOpts.Name, CreateOpts.ResourceGroupName,
		CreateOpts.CredRequestDir, CreateOpts.TargetDir)
	if err != nil {
		return err
	}

	return nil
}

func createServiceIDs(client ibmcloud.Client, accountID *string,
	name, resourceGroupName, credReqDir, targetDir string) error {

	var resourceGroupID string
	if resourceGroupName != "" {
		// Get the ID for the given resourceGroupName
		listResourceGroupsOptions := &resourcemanagerv2.ListResourceGroupsOptions{
			Name: &resourceGroupName,
		}
		resourceGroups, _, err := client.ListResourceGroups(listResourceGroupsOptions)
		if err != nil {
			return errors.Wrapf(err, "Failed to list resource groups for the name: %s", resourceGroupName)
		}

		if len(resourceGroups.Resources) == 0 {
			return errors.Errorf("Resource group %s not found", resourceGroupName)
		}

		resourceGroupID = *resourceGroups.Resources[0].ID
	}

	// Process directory
	credReqs, err := getListOfCredentialsRequests(credReqDir)
	if err != nil {
		return errors.Wrap(err, "Failed to process files containing CredentialsRequests")
	}

	var serviceIDs []*ServiceID

	undo := func() {
		for _, serviceID := range serviceIDs {
			serviceID.UnDo(targetDir)
		}
	}

	for _, cr := range credReqs {
		serviceID := NewServiceID(client, name, *accountID, resourceGroupID, cr)
		serviceIDs = append(serviceIDs, serviceID)
	}

	for _, serviceID := range serviceIDs {
		if err := serviceID.Validate(); err != nil {
			return errors.Wrap(err, "Failed to validate the serviceID")
		}
	}

	for _, serviceID := range serviceIDs {
		if err := serviceID.Do(); err != nil {
			undo()
			return errors.Wrap(err, "Failed to process the serviceID")
		}
	}

	for _, serviceID := range serviceIDs {
		if err := serviceID.Dump(targetDir); err != nil {
			undo()
			return errors.Wrap(err, "Failed to generate the secrets for serviceIDs")
		}
	}

	return nil
}

func getListOfCredentialsRequests(dir string) ([]*credreqv1.CredentialsRequest, error) {
	var credRequests []*credreqv1.CredentialsRequest
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		err = func() error {
			f, err := os.Open(filepath.Join(dir, file.Name()))
			if err != nil {
				return errors.Wrap(err, "Failed to open file")
			}
			defer f.Close()
			decoder := yaml.NewYAMLOrJSONDecoder(f, 4096)
			for {
				cr := &credreqv1.CredentialsRequest{}
				if err := decoder.Decode(cr); err != nil {
					if err == io.EOF {
						break
					}
					return errors.Wrap(err, "Failed to decode to CredentialsRequest")
				}
				credRequests = append(credRequests, cr)
			}
			return nil
		}()

		if err != nil {
			return nil, err
		}
	}

	return credRequests, nil
}

// initEnvForCreateServiceIDCmd will ensure the destination directory is ready to receive the generated
// files, and will create the directory if necessary.
func initEnvForCreateServiceIDCmd(cmd *cobra.Command, args []string) {
	if CreateOpts.TargetDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current directory: %s", err)
		}

		CreateOpts.TargetDir = pwd
	}

	fPath, err := filepath.Abs(CreateOpts.TargetDir)
	if err != nil {
		log.Fatalf("Failed to resolve full path: %s", err)
	}

	// create target dir if necessary
	err = provisioning.EnsureDir(fPath)
	if err != nil {
		log.Fatalf("failed to create target directory at %s", fPath)
	}

	// create manifests dir if necessary
	manifestsDir := filepath.Join(fPath, manifestsDirName)
	err = provisioning.EnsureDir(manifestsDir)
	if err != nil {
		log.Fatalf("failed to create manifests directory at %s", manifestsDir)
	}
}
