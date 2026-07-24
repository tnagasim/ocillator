package main

import (
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "ocillator",
		Short:   "OCI image distribution tool for air-gapped environments",
		Version: version,
	}

	root.AddCommand(newSyncCmd(), newReleaseCmd(), newDeployCmd())
	return root
}

func newSyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Sync OCI images to local layout",
		RunE:  func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func newReleaseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "release",
		Short: "Manage release history",
		RunE:  func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func newDeployCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy OCI images to target environment",
		RunE:  func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
