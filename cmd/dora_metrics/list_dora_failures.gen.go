package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ListDORAFailuresCmd = &cobra.Command{
	Use: "list-dora-failures [payload]",

	Short: "Get a list of failure events",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.DORAListFailuresRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORAFailures(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-dora-failures")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListDORAFailuresCmd)
}
