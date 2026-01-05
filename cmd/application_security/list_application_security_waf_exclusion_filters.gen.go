package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListApplicationSecurityWafExclusionFiltersCmd = &cobra.Command{
	Use:     "list-application-security-waf-exclusion-filters",
	Aliases: []string{"list-waf-exclusion-filters"},
	Short:   "List all WAF exclusion filters",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.ListApplicationSecurityWafExclusionFilters(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-application-security-waf-exclusion-filters")

		cmd.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWafExclusionFiltersCmd)
}
