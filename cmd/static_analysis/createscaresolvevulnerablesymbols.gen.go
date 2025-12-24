package static_analysis

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use:   "createscaresolvevulnerablesymbols",
	Short: "POST request to resolve vulnerable symbols",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/static-analysis-sca/vulnerabilities/resolve-vulnerable-symbols")
		fmt.Println("OperationID: CreateSCAResolveVulnerableSymbols")
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
