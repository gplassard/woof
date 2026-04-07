package widgets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteWidgetCmd = &cobra.Command{
	Use:     "delete-widget [experience_type] [uuid]",
	Aliases: []string{"delete"},
	Short:   "Delete a widget",
	Long: `Delete a widget
Documentation: https://docs.datadoghq.com/api/latest/widgets/#delete-widget`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewWidgetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteWidget(client.NewContext(apiKey, appKey, site), datadogV2.WidgetExperienceType(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to delete-widget")

	},
}

func init() {

	Cmd.AddCommand(DeleteWidgetCmd)
}
