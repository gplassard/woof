package application_security

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "updateapplicationsecuritywafcustomrule [custom_rule_id]",
	Short: "Update a WAF Custom Rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.UpdateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationSecurityWafCustomRuleUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateapplicationsecuritywafcustomrule: %v", err)
		}

		cmdutil.PrintJSON(res, "application_security")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafCustomRuleCmd)
}
