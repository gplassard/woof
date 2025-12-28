package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "update_application_security_waf_exclusion_filter [exclusion_filter_id]",
	Short: "Update a WAF exclusion filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.UpdateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationSecurityWafExclusionFilterUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_application_security_waf_exclusion_filter: %v", err)
		}

		cmdutil.PrintJSON(res, "exclusion_filter")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafExclusionFilterCmd)
}
