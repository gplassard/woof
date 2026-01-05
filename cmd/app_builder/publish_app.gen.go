package app_builder

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var PublishAppCmd = &cobra.Command{
	Use: "publish-app [app_id]",

	Short: "Publish App",
	Long: `Publish App
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#publish-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PublishAppResponse
		var err error

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err = api.PublishApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to publish-app")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {
	Cmd.AddCommand(PublishAppCmd)
}
