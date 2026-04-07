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

var UpdateWidgetCmd = &cobra.Command{
	Use:     "update-widget [experience_type] [uuid]",
	Aliases: []string{"update"},
	Short:   "Update a widget",
	Long: `Update a widget
Documentation: https://docs.datadoghq.com/api/latest/widgets/#update-widget`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WidgetResponse
		var err error

		var body datadogV2.CreateOrUpdateWidgetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewWidgetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateWidget(client.NewContext(apiKey, appKey, site), datadogV2.WidgetExperienceType(args[0]), uuid.MustParse(args[1]), body)
		cmdutil.HandleError(err, "failed to update-widget")

		fmt.Println(cmdutil.FormatJSON(res, "widgets"))
	},
}

func init() {

	UpdateWidgetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateWidgetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateWidgetCmd)
}
