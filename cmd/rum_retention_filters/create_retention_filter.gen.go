package rum_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRetentionFilterCmd = &cobra.Command{
	Use: "create-retention-filter [app_id]",

	Short: "Create a RUM retention filter",
	Long: `Create a RUM retention filter
Documentation: https://docs.datadoghq.com/api/latest/rum-retention-filters/#create-retention-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumRetentionFilterResponse
		var err error

		var body datadogV2.RumRetentionFilterCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-retention-filter")

		fmt.Println(cmdutil.FormatJSON(res, "retention_filter"))
	},
}

func init() {

	CreateRetentionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRetentionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRetentionFilterCmd)
}
