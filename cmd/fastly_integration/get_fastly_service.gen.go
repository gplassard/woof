package fastly_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFastlyServiceCmd = &cobra.Command{
	Use: "get-fastly-service [account_id] [service_id]",

	Short: "Get Fastly service",
	Long: `Get Fastly service
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#get-fastly-service`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyServiceResponse
		var err error

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-fastly-service")

		fmt.Println(cmdutil.FormatJSON(res, "fastly_service"))
	},
}

func init() {

	Cmd.AddCommand(GetFastlyServiceCmd)
}
