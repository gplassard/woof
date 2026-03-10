package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentPostmortemTemplatesCmd = &cobra.Command{
	Use:     "list-incident-postmortem-templates",
	Aliases: []string{"list-postmortem-templates"},
	Short:   "List postmortem templates",
	Long: `List postmortem templates
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-postmortem-templates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PostmortemTemplatesResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentPostmortemTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-postmortem-templates")

		fmt.Println(cmdutil.FormatJSON(res, "postmortem_template"))
	},
}

func init() {

	Cmd.AddCommand(ListIncidentPostmortemTemplatesCmd)
}
