package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLogsCustomDestinationCmd = &cobra.Command{
	Use:     "create-logs-custom-destination",
	Aliases: []string{"create"},
	Short:   "Create a custom destination",
	Long: `Create a custom destination
Documentation: https://docs.datadoghq.com/api/latest/logs-custom-destinations/#create-logs-custom-destination`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationResponse
		var err error

		var body datadogV2.CustomDestinationCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err = api.CreateLogsCustomDestination(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-custom-destination")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {

	CreateLogsCustomDestinationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLogsCustomDestinationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLogsCustomDestinationCmd)
}
