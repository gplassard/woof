package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDORAFailureCmd = &cobra.Command{
	Use: "create-dora-failure",

	Short: "Send a failure event",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORAFailure(client.NewContext(apiKey, appKey, site), datadogV2.DORAFailureRequest{})
		cmdutil.HandleError(err, "failed to create-dora-failure")

		cmd.Println(cmdutil.FormatJSON(res, "dora_failure"))
	},
}

func init() {
	Cmd.AddCommand(CreateDORAFailureCmd)
}
