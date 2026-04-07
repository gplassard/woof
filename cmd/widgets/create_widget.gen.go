package widgets

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateWidgetCmd = &cobra.Command{
	Use:     "create-widget [experience_type]",
	Aliases: []string{"create"},
	Short:   "Create a widget",
	Long: `Create a widget
Documentation: https://docs.datadoghq.com/api/latest/widgets/#create-widget`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WidgetResponse
		var err error

		var body datadogV2.CreateOrUpdateWidgetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewWidgetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateWidget(client.NewContext(apiKey, appKey, site), datadogV2.WidgetExperienceType(args[0]), body)
		cmdutil.HandleError(err, "failed to create-widget")

		fmt.Println(cmdutil.FormatJSON(res, "widgets"))
	},
}

func init() {

	CreateWidgetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateWidgetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateWidgetCmd)
}
