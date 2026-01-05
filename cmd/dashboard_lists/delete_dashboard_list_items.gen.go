package dashboard_lists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var DeleteDashboardListItemsCmd = &cobra.Command{
	Use:     "delete-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"delete-items"},
	Short:   "Delete items from a dashboard list",
	Long: `Delete items from a dashboard list
Documentation: https://docs.datadoghq.com/api/latest/dashboard-lists/#delete-dashboard-list-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DashboardListDeleteItemsResponse
		var err error

		var body datadogV2.DashboardListDeleteItemsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err = api.DeleteDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to delete-dashboard-list-items")

		cmd.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {

	DeleteDashboardListItemsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteDashboardListItemsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteDashboardListItemsCmd)
}
