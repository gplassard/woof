package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFastlyServiceCmd = &cobra.Command{
	Use: "update-fastly-service [account_id] [service_id]",

	Short: "Update Fastly service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.FastlyServiceRequest{})
		cmdutil.HandleError(err, "failed to update-fastly-service")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyServiceCmd)
}
