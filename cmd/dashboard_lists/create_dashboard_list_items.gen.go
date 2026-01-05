package dashboard_lists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"encoding/json"
	"github.com/spf13/cobra"
	"strconv"
)

var CreateDashboardListItemsCmd = &cobra.Command{
	Use:     "create-dashboard-list-items [dashboard_list_id] [payload]",
	Aliases: []string{"create-items"},
	Short:   "Add Items to a Dashboard List",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.DashboardListAddItemsRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err := api.CreateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to create-dashboard-list-items")

		cmd.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {
	Cmd.AddCommand(CreateDashboardListItemsCmd)
}
