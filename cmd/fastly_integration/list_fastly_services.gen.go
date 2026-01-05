package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFastlyServicesCmd = &cobra.Command{
	Use: "list-fastly-services [account_id]",

	Short: "List Fastly services",
	Long: `List Fastly services
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#list-fastly-services`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyServicesResponse
		var err error

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err = api.ListFastlyServices(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-fastly-services")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {
	Cmd.AddCommand(ListFastlyServicesCmd)
}
