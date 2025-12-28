package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "create-application-security-waf-exclusion-filter",
	Short: "Create a WAF exclusion filter",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.CreateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), datadogV2.ApplicationSecurityWafExclusionFilterCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-application-security-waf-exclusion-filter: %v", err)
		}

		cmdutil.PrintJSON(res, "exclusion_filter")
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafExclusionFilterCmd)
}
