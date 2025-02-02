package gcp

import (
	"github.com/spf13/cobra"

	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning"
)

type options struct {
	TargetDir      string
	PublicKeyPath  string
	Region         string
	Name           string
	CredRequestDir string
	DryRun         bool
}

// NewGCPCmd implements the "gcp" subcommand for the credentials provisioning
func NewGCPCmd() *cobra.Command {
	gcpCmd := &cobra.Command{
		Use:   "gcp",
		Short: "Manage credentials objects for Google cloud",
		Long:  "Creating/updating/deleting cloud credentials objects for Google cloud",
	}

	gcpCmd.AddCommand(provisioning.NewCreateKeyPairCmd())

	return gcpCmd
}
