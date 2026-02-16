package action_connection

import (
	"fmt"
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

		optionalParams := datadogV2.NewListAppKeyRegistrationsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAppKeyRegistrations(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-app-key-registrations")

		fmt.Println(cmdutil.FormatJSON(res, "app_key_registration"))
	},
}

func init() {

	ListAppKeyRegistrationsCmd.Flags().Int64("page-size", 0, "The number of App Key Registrations to return per page.")

	ListAppKeyRegistrationsCmd.Flags().Int64("page-number", 0, "The page number to return.")

	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
