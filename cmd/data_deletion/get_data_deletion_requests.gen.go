package data_deletion

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDataDeletionRequestsCmd = &cobra.Command{
	Use:     "get-data-deletion-requests",
	Aliases: []string{"get-requests"},
	Short:   "Gets a list of data deletion requests",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		res, _, err := api.GetDataDeletionRequests(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-data-deletion-requests")

		cmd.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {
	Cmd.AddCommand(GetDataDeletionRequestsCmd)
}
