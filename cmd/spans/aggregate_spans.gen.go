package spans

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AggregateSpansCmd = &cobra.Command{
	Use:     "aggregate-spans [payload]",
	Aliases: []string{"aggregate"},
	Short:   "Aggregate spans",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SpansAggregateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.AggregateSpans(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-spans")

		cmd.Println(cmdutil.FormatJSON(res, "bucket"))
	},
}

func init() {
	Cmd.AddCommand(AggregateSpansCmd)
}
