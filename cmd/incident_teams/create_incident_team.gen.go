package incident_teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentTeamCmd = &cobra.Command{
	Use:     "create-incident-team",
	Aliases: []string{"create"},
	Short:   "Create a new incident team",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTeam(client.NewContext(apiKey, appKey, site), datadogV2.IncidentTeamCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-team")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTeamCmd)
}
