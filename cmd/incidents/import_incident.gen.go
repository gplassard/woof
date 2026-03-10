package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ImportIncidentCmd = &cobra.Command{
	Use:     "import-incident",
	Aliases: []string{"import"},
	Short:   "Import an incident",
	Long: `Import an incident
Documentation: https://docs.datadoghq.com/api/latest/incidents/#import-incident`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentImportResponse
		var err error

		var body datadogV2.IncidentImportRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ImportIncident(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to import-incident")

		fmt.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	ImportIncidentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ImportIncidentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ImportIncidentCmd)
}
