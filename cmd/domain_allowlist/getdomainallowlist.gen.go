package domain_allowlist

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDomainAllowlistCmd = &cobra.Command{
	Use:   "getdomainallowlist",
	Short: "Get Domain Allowlist",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/domain_allowlist")
		fmt.Println("OperationID: GetDomainAllowlist")
	},
}

func init() {
	Cmd.AddCommand(GetDomainAllowlistCmd)
}
