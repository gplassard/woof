package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecretsRulesCmd = &cobra.Command{
	Use:   "getsecretsrules",
	Short: "Returns a list of Secrets rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/static-analysis/secrets/rules")
		fmt.Println("OperationID: GetSecretsRules")
	},
}

func init() {
	Cmd.AddCommand(GetSecretsRulesCmd)
}
