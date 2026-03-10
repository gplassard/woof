package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentPostmortemTemplateCmd = &cobra.Command{
	Use:     "get-incident-postmortem-template [template_id]",
	Aliases: []string{"get-postmortem-template"},
	Short:   "Get postmortem template",
	Long: `Get postmortem template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-incident-postmortem-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PostmortemTemplateResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetIncidentPostmortemTemplate(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-incident-postmortem-template")

		fmt.Println(cmdutil.FormatJSON(res, "postmortem_template"))
	},
}

func init() {

	Cmd.AddCommand(GetIncidentPostmortemTemplateCmd)
}
