package events

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateEventCmd = &cobra.Command{
	Use:     "create-event",
	Aliases: []string{"create"},
	Short:   "Post an event",
	Long: `Post an event
Documentation: https://docs.datadoghq.com/api/latest/events/#create-event`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventCreateResponsePayload
		var err error

		var body datadogV2.EventCreateRequestPayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateEvent(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-event")

		fmt.Println(cmdutil.FormatJSON(res, "event"))
	},
}

func init() {

	CreateEventCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateEventCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateEventCmd)
}
