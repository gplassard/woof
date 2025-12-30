package fastly_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFastlyServiceCmd = &cobra.Command{
	Use: "create-fastly-service [account_id]",

	Short: "Add Fastly service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateFastlyService(client.NewContext(apiKey, appKey, site), args[0], datadogV2.FastlyServiceRequest{})
		cmdutil.HandleError(err, "failed to create-fastly-service")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyServiceCmd)
}
