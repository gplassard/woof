package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteFastlyServiceCmd = &cobra.Command{
	Use: "delete-fastly-service [account_id] [service_id]",

	Short: "Delete Fastly service",
	Long: `Delete Fastly service
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#delete-fastly-service`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-fastly-service")

	},
}

func init() {

	Cmd.AddCommand(DeleteFastlyServiceCmd)
}
