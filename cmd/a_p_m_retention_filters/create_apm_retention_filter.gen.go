package a_p_m_retention_filters

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateApmRetentionFilterCmd = &cobra.Command{
	Use: "create-apm-retention-filter",

	Short: "Create a retention filter",
	Long: `Create a retention filter
Documentation: https://docs.datadoghq.com/api/latest/a-p-m-retention-filters/#create-apm-retention-filter`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RetentionFilterCreateResponse
		var err error

		var body datadogV2.RetentionFilterCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateApmRetentionFilter(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-apm-retention-filter")

		fmt.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {

	CreateApmRetentionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateApmRetentionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateApmRetentionFilterCmd)
}
