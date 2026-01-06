package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAppKeyRegistrationsCmd = &cobra.Command{
	Use: "list-app-key-registrations",

	Short: "List App Key Registrations",
	Long: `List App Key Registrations
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#list-app-key-registrations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListAppKeyRegistrationsResponse
		var err error

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAppKeyRegistrations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-app-key-registrations")

		cmd.Println(cmdutil.FormatJSON(res, "app_key_registration"))
	},
}

func init() {

	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
