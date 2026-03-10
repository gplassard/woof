package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateGlobalIncidentHandleCmd = &cobra.Command{
	Use:     "create-global-incident-handle",
	Aliases: []string{"create-global-handle"},
	Short:   "Create global incident handle",
	Long: `Create global incident handle
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-global-incident-handle`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentHandleResponse
		var err error

		var body datadogV2.IncidentHandleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateGlobalIncidentHandle(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-global-incident-handle")

		fmt.Println(cmdutil.FormatJSON(res, "incidents_handles"))
	},
}

func init() {

	CreateGlobalIncidentHandleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateGlobalIncidentHandleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateGlobalIncidentHandleCmd)
}
