package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ConvertSecurityMonitoringRuleFromJSONToTerraformCmd = &cobra.Command{
	Use:   "convertsecuritymonitoringrulefromjsontoterraform",
	Short: "Convert a rule from JSON to Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/rules/convert")
		fmt.Println("OperationID: ConvertSecurityMonitoringRuleFromJSONToTerraform")
	},
}

func init() {
	Cmd.AddCommand(ConvertSecurityMonitoringRuleFromJSONToTerraformCmd)
}
