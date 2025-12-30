package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentsCmd = &cobra.Command{
	Use:     "list-incidents",
	Aliases: []string{"list"},
	Short:   "Get a list of incidents",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentsCmd)
}
