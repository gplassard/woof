package application_security

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "createapplicationsecuritywafcustomrule",
	Short: "Create a WAF custom rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.CreateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), datadogV2.ApplicationSecurityWafCustomRuleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createapplicationsecuritywafcustomrule: %v", err)
		}

		cmdutil.PrintJSON(res, "application_security")
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafCustomRuleCmd)
}
