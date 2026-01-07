package app_builder

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateAppCmd = &cobra.Command{
	Use: "update-app [app_id]",

	Short: "Update App",
	Long: `Update App
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#update-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateAppResponse
		var err error

		var body datadogV2.UpdateAppRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-app")

		fmt.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {

	UpdateAppCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAppCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAppCmd)
}
