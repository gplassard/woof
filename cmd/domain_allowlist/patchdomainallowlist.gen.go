package domain_allowlist

import (
	"fmt"
	"github.com/spf13/cobra"
)

var PatchDomainAllowlistCmd = &cobra.Command{
	Use:   "patchdomainallowlist",
	Short: "Sets Domain Allowlist",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/domain_allowlist")
		fmt.Println("OperationID: PatchDomainAllowlist")
	},
}

func init() {
	Cmd.AddCommand(PatchDomainAllowlistCmd)
}
