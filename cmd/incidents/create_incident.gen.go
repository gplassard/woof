package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentCmd = &cobra.Command{
	Use:     "create-incident",
	Aliases: []string{"create"},
	Short:   "Create an incident",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncident(client.NewContext(apiKey, appKey, site), datadogV2.IncidentCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentCmd)
}
