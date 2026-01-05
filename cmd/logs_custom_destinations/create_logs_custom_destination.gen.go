package logs_custom_destinations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateLogsCustomDestinationCmd = &cobra.Command{
	Use:     "create-logs-custom-destination [payload]",
	Aliases: []string{"create"},
	Short:   "Create a custom destination",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomDestinationResponse
		var err error

		var body datadogV2.CustomDestinationCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsCustomDestinationsApi(client.NewAPIClient())
		res, _, err = api.CreateLogsCustomDestination(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-custom-destination")

		cmd.Println(cmdutil.FormatJSON(res, "custom_destination"))
	},
}

func init() {
	Cmd.AddCommand(CreateLogsCustomDestinationCmd)
}
