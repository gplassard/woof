package opsgenie_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOpsgenieServiceCmd = &cobra.Command{
	Use: "get-opsgenie-service [integration_service_id]",

	Short: "Get a single service object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetOpsgenieService(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-opsgenie-service")

		cmd.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {
	Cmd.AddCommand(GetOpsgenieServiceCmd)
}
