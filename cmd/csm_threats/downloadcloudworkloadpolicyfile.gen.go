package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DownloadCloudWorkloadPolicyFileCmd = &cobra.Command{
	Use:   "downloadcloudworkloadpolicyfile",
	Short: "Download the Workload Protection policy (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/cloud_workload/policy/download")
		fmt.Println("OperationID: DownloadCloudWorkloadPolicyFile")
	},
}

func init() {
	Cmd.AddCommand(DownloadCloudWorkloadPolicyFileCmd)
}
