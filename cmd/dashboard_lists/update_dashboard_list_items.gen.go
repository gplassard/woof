package dashboard_lists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var UpdateDashboardListItemsCmd = &cobra.Command{
	Use:     "update-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"update-items"},
	Short:   "Update items of a dashboard list",
	Long: `Update items of a dashboard list
Documentation: https://docs.datadoghq.com/api/latest/dashboard-lists/#update-dashboard-list-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DashboardListUpdateItemsResponse
		var err error

		var body datadogV2.DashboardListUpdateItemsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err = api.UpdateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-dashboard-list-items")

		cmd.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {

	UpdateDashboardListItemsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDashboardListItemsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDashboardListItemsCmd)
}
