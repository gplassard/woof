package apm_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderApmRetentionFiltersCmd = &cobra.Command{
	Use:     "reorder-apm-retention-filters",
	Aliases: []string{"reorder"},
	Short:   "Re-order retention filters",
	Long: `Re-order retention filters
Documentation: https://docs.datadoghq.com/api/latest/apm-retention-filters/#reorder-apm-retention-filters`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ReorderRetentionFiltersRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ReorderApmRetentionFilters(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-apm-retention-filters")

	},
}

func init() {

	ReorderApmRetentionFiltersCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ReorderApmRetentionFiltersCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ReorderApmRetentionFiltersCmd)
}
