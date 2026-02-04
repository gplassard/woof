package fastly_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFastlyAccountCmd = &cobra.Command{
	Use: "get-fastly-account [account_id]",

	Short: "Get Fastly account",
	Long: `Get Fastly account
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#get-fastly-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyAccountResponse
		var err error

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFastlyAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fastly-account")

		fmt.Println(cmdutil.FormatJSON(res, "fastly_account"))
	},
}

func init() {

	Cmd.AddCommand(GetFastlyAccountCmd)
}
