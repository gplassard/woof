package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DownloadCSMThreatsPolicyCmd = &cobra.Command{
	Use:   "downloadcsmthreatspolicy",
	Short: "Download the Workload Protection policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/cws/policy/download")
		fmt.Println("OperationID: DownloadCSMThreatsPolicy")
	},
}

func init() {
	Cmd.AddCommand(DownloadCSMThreatsPolicyCmd)
}
