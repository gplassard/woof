package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var QueryEventFilteredUsersCmd = &cobra.Command{
	Use: "query-event-filtered-users [payload]",

	Short: "Query event filtered users",
	Long: `Query event filtered users
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#query-event-filtered-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.QueryResponse
		var err error

		var body datadogV2.QueryEventFilteredUsersRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err = api.QueryEventFilteredUsers(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-event-filtered-users")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryEventFilteredUsersCmd)
}
