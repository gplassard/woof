package logs

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AggregateLogsCmd = &cobra.Command{
	Use:     "aggregate-logs [payload]",
	Aliases: []string{"aggregate"},
	Short:   "Aggregate events",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.LogsAggregateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		res, _, err := api.AggregateLogs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-logs")

		cmd.Println(cmdutil.FormatJSON(res, "logs"))
	},
}

func init() {
	Cmd.AddCommand(AggregateLogsCmd)
}
