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

var UnpublishAppCmd = &cobra.Command{
	Use: "unpublish-app [app_id]",

	Short: "Unpublish App",
	Long: `Unpublish App
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#unpublish-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UnpublishAppResponse
		var err error

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UnpublishApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to unpublish-app")

		fmt.Println(cmdutil.FormatJSON(res, "app_builder"))
	},
}

func init() {

	Cmd.AddCommand(UnpublishAppCmd)
}
