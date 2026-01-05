package opsgenie_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteOpsgenieServiceCmd = &cobra.Command{
	Use: "delete-opsgenie-service [integration_service_id]",

	Short: "Delete a single service object",
	Long: `Delete a single service object
Documentation: https://docs.datadoghq.com/api/latest/opsgenie-integration/#delete-opsgenie-service`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		_, err = api.DeleteOpsgenieService(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-opsgenie-service")

	},
}

func init() {
	Cmd.AddCommand(DeleteOpsgenieServiceCmd)
}
