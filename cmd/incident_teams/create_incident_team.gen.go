package incident_teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentTeamCmd = &cobra.Command{
	Use:     "create-incident-team [payload]",
	Aliases: []string{"create"},
	Short:   "Create a new incident team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.IncidentTeamCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTeam(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-team")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTeamCmd)
}
