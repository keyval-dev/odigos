/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/keyval-dev/odigos/cli/cmd/resources"
	"github.com/keyval-dev/odigos/cli/pkg/kube"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	OdigosVersion string
	OdigosCommit  string
	OdigosDate    string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print odigos version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Odigos Cli Version: version.Info{Version:'%s', GitCommit:'%s', BuildDate:'%s'}\n", OdigosVersion, OdigosCommit, OdigosDate)
		printOdigosClusterVersion(cmd)
	},
}

func printOdigosClusterVersion(cmd *cobra.Command) {
	client, err := kube.CreateClient(cmd)
	if err != nil {
		return
	}
	ctx := cmd.Context()

	ns, err := resources.GetOdigosNamespace(client, ctx)
	if err != nil {
		if resources.IsErrNoOdigosNamespaceFound(err) {
			fmt.Println("Odigos is NOT yet installed in the current cluster")
		} else {
			fmt.Println("Error detecting Odigos version in the current cluster")
		}
		return
	}

	cm, err := client.CoreV1().ConfigMaps(ns).Get(ctx, resources.OdigosDeploymentConfigMapName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("Error detecting Odigos version in the current cluster")
		return
	}

	odigosVersion, ok := cm.Data["ODIGOS_VERSION"]
	if !ok || odigosVersion == "" {
		fmt.Println("Error detecting Odigos version in the current cluster")
		return
	}
	fmt.Printf("Odigos Version (in cluster): version.Info{Version:'%s'}\n", odigosVersion)
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
