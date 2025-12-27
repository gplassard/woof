package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "getapplicationsecuritywafexclusionfilter [exclusion_filter_id]",
	Short: "Get a WAF exclusion filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.GetApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getapplicationsecuritywafexclusionfilter: %v", err)
		}

		cmdutil.PrintJSON(res, "application_security")
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafExclusionFilterCmd)
}
