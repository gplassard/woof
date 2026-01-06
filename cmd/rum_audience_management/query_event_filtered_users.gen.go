package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryEventFilteredUsersCmd = &cobra.Command{
	Use: "query-event-filtered-users",

	Short: "Query event filtered users",
	Long: `Query event filtered users
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#query-event-filtered-users`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.QueryResponse
		var err error

		var body datadogV2.QueryEventFilteredUsersRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.QueryEventFilteredUsers(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-event-filtered-users")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {

	QueryEventFilteredUsersCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryEventFilteredUsersCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryEventFilteredUsersCmd)
}
