package fastly_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFastlyServiceCmd = &cobra.Command{
	Use: "create-fastly-service [account_id]",

	Short: "Add Fastly service",
	Long: `Add Fastly service
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#create-fastly-service`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyServiceResponse
		var err error

		var body datadogV2.FastlyServiceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateFastlyService(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-fastly-service")

		fmt.Println(cmdutil.FormatJSON(res, "fastly_service"))
	},
}

func init() {

	CreateFastlyServiceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateFastlyServiceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateFastlyServiceCmd)
}
