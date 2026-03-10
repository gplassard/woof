package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentPostmortemTemplateCmd = &cobra.Command{
	Use:     "create-incident-postmortem-template",
	Aliases: []string{"create-postmortem-template"},
	Short:   "Create postmortem template",
	Long: `Create postmortem template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-postmortem-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PostmortemTemplateResponse
		var err error

		var body datadogV2.PostmortemTemplateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentPostmortemTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-postmortem-template")

		fmt.Println(cmdutil.FormatJSON(res, "postmortem_template"))
	},
}

func init() {

	CreateIncidentPostmortemTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentPostmortemTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentPostmortemTemplateCmd)
}
