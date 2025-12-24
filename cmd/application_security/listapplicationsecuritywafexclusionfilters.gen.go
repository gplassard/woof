package application_security

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListApplicationSecurityWafExclusionFiltersCmd = &cobra.Command{
	Use:   "listapplicationsecuritywafexclusionfilters",
	Short: "List all WAF exclusion filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.ListApplicationSecurityWafExclusionFilters(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listapplicationsecuritywafexclusionfilters: %v", err)
		}

		cmdutil.PrintJSON(res, "application_security")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWafExclusionFiltersCmd)
}
