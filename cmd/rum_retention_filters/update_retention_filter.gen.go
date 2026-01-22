package rum_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRetentionFilterCmd = &cobra.Command{
	Use: "update-retention-filter [app_id] [rf_id]",

	Short: "Update a RUM retention filter",
	Long: `Update a RUM retention filter
Documentation: https://docs.datadoghq.com/api/latest/rum-retention-filters/#update-retention-filter`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumRetentionFilterResponse
		var err error

		var body datadogV2.RumRetentionFilterUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-retention-filter")

		fmt.Println(cmdutil.FormatJSON(res, "retention_filter"))
	},
}

func init() {

	UpdateRetentionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRetentionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRetentionFilterCmd)
}
