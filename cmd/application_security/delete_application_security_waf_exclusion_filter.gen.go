package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "delete-application-security-waf-exclusion-filter [exclusion_filter_id]",
	Short: "Delete a WAF exclusion filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		_, err := api.DeleteApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-application-security-waf-exclusion-filter: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationSecurityWafExclusionFilterCmd)
}
