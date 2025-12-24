package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ConvertSecurityMonitoringRuleFromJSONToTerraformCmd = &cobra.Command{
	Use:   "convertsecuritymonitoringrulefromjsontoterraform",
	Short: "Convert a rule from JSON to Terraform",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ConvertSecurityMonitoringRuleFromJSONToTerraform(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleConvertPayload{})
		if err != nil {
			log.Fatalf("failed to convertsecuritymonitoringrulefromjsontoterraform: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ConvertSecurityMonitoringRuleFromJSONToTerraformCmd)
}
