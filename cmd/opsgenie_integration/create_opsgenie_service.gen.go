package opsgenie_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOpsgenieServiceCmd = &cobra.Command{
	Use: "create-opsgenie-service [payload]",

	Short: "Create a new service object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OpsgenieServiceResponse
		var err error

		var body datadogV2.OpsgenieServiceCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err = api.CreateOpsgenieService(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-opsgenie-service")

		cmd.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {
	Cmd.AddCommand(CreateOpsgenieServiceCmd)
}
