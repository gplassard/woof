package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateFastlyServiceCmd = &cobra.Command{
	Use: "update-fastly-service [account_id] [service_id] [payload]",

	Short: "Update Fastly service",
	Long: `Update Fastly service
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#update-fastly-service`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyServiceResponse
		var err error

		var body datadogV2.FastlyServiceRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-fastly-service")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyServiceCmd)
}
