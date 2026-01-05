package apm_retention_filters

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListApmRetentionFiltersCmd = &cobra.Command{
	Use:     "list-apm-retention-filters",
	Aliases: []string{"list"},
	Short:   "List all APM retention filters",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.ListApmRetentionFilters(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-apm-retention-filters")

		cmd.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {
	Cmd.AddCommand(ListApmRetentionFiltersCmd)
}
