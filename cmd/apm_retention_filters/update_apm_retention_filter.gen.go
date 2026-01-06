package apm_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateApmRetentionFilterCmd = &cobra.Command{
	Use:     "update-apm-retention-filter [filter_id]",
	Aliases: []string{"update"},
	Short:   "Update a retention filter",
	Long: `Update a retention filter
Documentation: https://docs.datadoghq.com/api/latest/apm-retention-filters/#update-apm-retention-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RetentionFilterResponse
		var err error

		var body datadogV2.RetentionFilterUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateApmRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-apm-retention-filter")

		cmd.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {

	UpdateApmRetentionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateApmRetentionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateApmRetentionFilterCmd)
}
