package logs

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AggregateLogsCmd = &cobra.Command{
	Use:     "aggregate-logs",
	Aliases: []string{"aggregate"},
	Short:   "Aggregate events",
	Long: `Aggregate events
Documentation: https://docs.datadoghq.com/api/latest/logs/#aggregate-logs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsAggregateResponse
		var err error

		var body datadogV2.LogsAggregateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AggregateLogs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-logs")

		fmt.Println(cmdutil.FormatJSON(res, "logs"))
	},
}

func init() {

	AggregateLogsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AggregateLogsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AggregateLogsCmd)
}
