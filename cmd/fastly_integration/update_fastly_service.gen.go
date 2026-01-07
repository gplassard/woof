package fastly_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFastlyServiceCmd = &cobra.Command{
	Use: "update-fastly-service [account_id] [service_id]",

	Short: "Update Fastly service",
	Long: `Update Fastly service
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#update-fastly-service`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyServiceResponse
		var err error

		var body datadogV2.FastlyServiceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-fastly-service")

		fmt.Println(cmdutil.FormatJSON(res, "fastly-services"))
	},
}

func init() {

	UpdateFastlyServiceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFastlyServiceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFastlyServiceCmd)
}
