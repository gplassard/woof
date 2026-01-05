package rum_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRetentionFiltersCmd = &cobra.Command{
	Use: "list-retention-filters [app_id]",

	Short: "Get all RUM retention filters",
	Long: `Get all RUM retention filters
Documentation: https://docs.datadoghq.com/api/latest/rum-retention-filters/#list-retention-filters`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumRetentionFiltersResponse
		var err error

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err = api.ListRetentionFilters(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-retention-filters")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(ListRetentionFiltersCmd)
}
