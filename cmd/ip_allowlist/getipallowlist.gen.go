package ip_allowlist

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIPAllowlistCmd = &cobra.Command{
	Use:   "getipallowlist",
	Short: "Get IP Allowlist",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ip_allowlist")
		fmt.Println("OperationID: GetIPAllowlist")
	},
}

func init() {
	Cmd.AddCommand(GetIPAllowlistCmd)
}
