package a_p_m_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListApmRetentionFiltersCmd = &cobra.Command{
	Use: "list-apm-retention-filters",

	Short: "List all APM retention filters",
	Long: `List all APM retention filters
Documentation: https://docs.datadoghq.com/api/latest/a-p-m-retention-filters/#list-apm-retention-filters`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RetentionFiltersResponse
		var err error

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListApmRetentionFilters(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-apm-retention-filters")

		fmt.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {

	Cmd.AddCommand(ListApmRetentionFiltersCmd)
}
