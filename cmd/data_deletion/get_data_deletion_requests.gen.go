package data_deletion

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDataDeletionRequestsCmd = &cobra.Command{
	Use:     "get-data-deletion-requests",
	Aliases: []string{"get-requests"},
	Short:   "Gets a list of data deletion requests",
	Long: `Gets a list of data deletion requests
Documentation: https://docs.datadoghq.com/api/latest/data-deletion/#get-data-deletion-requests`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetDataDeletionsResponseBody
		var err error

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDataDeletionRequests(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-data-deletion-requests")

		cmd.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {

	Cmd.AddCommand(GetDataDeletionRequestsCmd)
}
