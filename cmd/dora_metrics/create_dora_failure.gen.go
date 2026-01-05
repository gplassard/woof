package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDORAFailureCmd = &cobra.Command{
	Use: "create-dora-failure [payload]",

	Short: "Send a failure event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAFailureResponse
		var err error

		var body datadogV2.DORAFailureRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err = api.CreateDORAFailure(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-dora-failure")

		cmd.Println(cmdutil.FormatJSON(res, "dora_failure"))
	},
}

func init() {
	Cmd.AddCommand(CreateDORAFailureCmd)
}
