package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDORAFailureCmd = &cobra.Command{
	Use: "delete-dora-failure [failure_id]",

	Short: "Delete a failure event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		_, err = api.DeleteDORAFailure(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-dora-failure")

	},
}

func init() {
	Cmd.AddCommand(DeleteDORAFailureCmd)
}
