package dashboard_lists

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var CreateDashboardListItemsCmd = &cobra.Command{
	Use:     "create-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"create-items"},
	Short:   "Add Items to a Dashboard List",
	Long: `Add Items to a Dashboard List
Documentation: https://docs.datadoghq.com/api/latest/dashboard-lists/#create-dashboard-list-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DashboardListAddItemsResponse
		var err error

		var body datadogV2.DashboardListAddItemsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to create-dashboard-list-items")

		fmt.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {

	CreateDashboardListItemsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDashboardListItemsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDashboardListItemsCmd)
}
