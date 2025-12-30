package fastly_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFastlyServicesCmd = &cobra.Command{
	Use: "list-fastly-services [account_id]",

	Short: "List Fastly services",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListFastlyServices(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-fastly-services")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {
	Cmd.AddCommand(ListFastlyServicesCmd)
}
