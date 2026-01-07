package data_deletion

import (
	"fmt"
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
	Long: `Cancels a data deletion request
Documentation: https://docs.datadoghq.com/api/latest/data-deletion/#cancel-data-deletion-request`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CancelDataDeletionResponseBody
		var err error

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CancelDataDeletionRequest(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to cancel-data-deletion-request")

		fmt.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {

	Cmd.AddCommand(CancelDataDeletionRequestCmd)
}
