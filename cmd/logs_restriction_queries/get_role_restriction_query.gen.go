package logs_restriction_queries

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRoleRestrictionQueryCmd = &cobra.Command{
	Use: "get-role-restriction-query [role_id]",

	Short: "Get restriction query for a given role",
	Long: `Get restriction query for a given role
Documentation: https://docs.datadoghq.com/api/latest/logs-restriction-queries/#get-role-restriction-query`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryListResponse
		var err error

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRoleRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-role-restriction-query")

		fmt.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {

	Cmd.AddCommand(GetRoleRestrictionQueryCmd)
}
