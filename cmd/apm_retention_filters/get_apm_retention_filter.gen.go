package apm_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetApmRetentionFilterCmd = &cobra.Command{
	Use:     "get-apm-retention-filter [filter_id]",
	Aliases: []string{"get"},
	Short:   "Get a given APM retention filter",
	Long: `Get a given APM retention filter
Documentation: https://docs.datadoghq.com/api/latest/apm-retention-filters/#get-apm-retention-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RetentionFilterResponse
		var err error

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetApmRetentionFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-apm-retention-filter")

		fmt.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {

	Cmd.AddCommand(GetApmRetentionFilterCmd)
}
