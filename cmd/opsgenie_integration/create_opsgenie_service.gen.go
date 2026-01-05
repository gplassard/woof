package opsgenie_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOpsgenieServiceCmd = &cobra.Command{
	Use: "create-opsgenie-service",

	Short: "Create a new service object",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateOpsgenieService(client.NewContext(apiKey, appKey, site), datadogV2.OpsgenieServiceCreateRequest{})
		cmdutil.HandleError(err, "failed to create-opsgenie-service")

		cmd.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {
	Cmd.AddCommand(CreateOpsgenieServiceCmd)
}
