package dashboard_lists

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var UpdateDashboardListItemsCmd = &cobra.Command{
	Use:   "updatedashboardlistitems [dashboard_list_id]",
	Short: "Update items of a dashboard list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err := api.UpdateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.DashboardListUpdateItemsRequest{})
		if err != nil {
			log.Fatalf("failed to updatedashboardlistitems: %v", err)
		}

		cmdutil.PrintJSON(res, "dashboard_lists")
	},
}

func init() {
	Cmd.AddCommand(UpdateDashboardListItemsCmd)
}
