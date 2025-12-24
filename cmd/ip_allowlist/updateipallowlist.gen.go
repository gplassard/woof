package ip_allowlist

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIPAllowlistCmd = &cobra.Command{
	Use:   "updateipallowlist",
	Short: "Update IP Allowlist",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/ip_allowlist")
		fmt.Println("OperationID: UpdateIPAllowlist")
	},
}

func init() {
	Cmd.AddCommand(UpdateIPAllowlistCmd)
}
