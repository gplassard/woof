package logs

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsCmd = &cobra.Command{
	Use:     "list-logs",
	Aliases: []string{"list"},
	Short:   "Search logs (POST)",
	Long: `Search logs (POST)
Documentation: https://docs.datadoghq.com/api/latest/logs/#list-logs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsListResponse
		var err error

		optionalParams := datadogV2.NewListLogsOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLogs(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-logs")

		cmd.Println(cmdutil.FormatJSON(res, "log"))
	},
}

func init() {

	ListLogsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ListLogsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ListLogsCmd)
}
