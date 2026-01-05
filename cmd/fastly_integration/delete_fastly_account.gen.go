package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteFastlyAccountCmd = &cobra.Command{
	Use: "delete-fastly-account [account_id]",

	Short: "Delete Fastly account",
	Long: `Delete Fastly account
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#delete-fastly-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		_, err = api.DeleteFastlyAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-fastly-account")

	},
}

func init() {

	Cmd.AddCommand(DeleteFastlyAccountCmd)
}
