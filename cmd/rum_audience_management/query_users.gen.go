package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryUsersCmd = &cobra.Command{
	Use: "query-users",

	Short: "Query users",
	Long: `Query users
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#query-users`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.QueryResponse
		var err error

		var body datadogV2.QueryUsersRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err = api.QueryUsers(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-users")

		cmd.Println(cmdutil.FormatJSON(res, "query_response"))
	},
}

func init() {

	QueryUsersCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryUsersCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryUsersCmd)
}
