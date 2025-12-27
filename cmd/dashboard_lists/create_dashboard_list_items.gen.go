package dashboard_lists

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var CreateDashboardListItemsCmd = &cobra.Command{
	Use:   "create_dashboard_list_items [dashboard_list_id]",
	Short: "Add Items to a Dashboard List",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err := api.CreateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.DashboardListAddItemsRequest{})
		if err != nil {
			log.Fatalf("failed to create_dashboard_list_items: %v", err)
		}

		cmdutil.PrintJSON(res, "dashboard_lists")
	},
}

func init() {
	Cmd.AddCommand(CreateDashboardListItemsCmd)
}
