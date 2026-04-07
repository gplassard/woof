package widgets

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetWidgetCmd = &cobra.Command{
	Use:     "get-widget [experience_type] [uuid]",
	Aliases: []string{"get"},
	Short:   "Get a widget",
	Long: `Get a widget
Documentation: https://docs.datadoghq.com/api/latest/widgets/#get-widget`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WidgetResponse
		var err error

		api := datadogV2.NewWidgetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetWidget(client.NewContext(apiKey, appKey, site), datadogV2.WidgetExperienceType(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to get-widget")

		fmt.Println(cmdutil.FormatJSON(res, "widgets"))
	},
}

func init() {

	Cmd.AddCommand(GetWidgetCmd)
}
