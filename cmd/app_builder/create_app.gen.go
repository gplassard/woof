package app_builder

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAppCmd = &cobra.Command{
	Use: "create-app",

	Short: "Create App",
	Long: `Create App
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#create-app`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateAppResponse
		var err error

		var body datadogV2.CreateAppRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateApp(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-app")

		fmt.Println(cmdutil.FormatJSON(res, "app"))
	},
}

func init() {

	CreateAppCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAppCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAppCmd)
}
