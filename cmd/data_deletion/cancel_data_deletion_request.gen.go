package data_deletion

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CancelDataDeletionRequestCmd = &cobra.Command{
	Use:     "cancel-data-deletion-request [id]",
	Aliases: []string{"cancel-request"},
	Short:   "Cancels a data deletion request",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		res, _, err := api.CancelDataDeletionRequest(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to cancel-data-deletion-request")

		cmd.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {
	Cmd.AddCommand(CancelDataDeletionRequestCmd)
}
