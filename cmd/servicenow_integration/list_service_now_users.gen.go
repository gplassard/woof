package servicenow_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var ListServiceNowUsersCmd = &cobra.Command{
	Use: "list-service-now-users [instance_id]",

	Short: "List ServiceNow users",
	Long: `List ServiceNow users
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#list-service-now-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowUsersResponse
		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceNowUsers(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to list-service-now-users")

		fmt.Println(cmdutil.FormatJSON(res, "servicenow_integration"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceNowUsersCmd)
}
