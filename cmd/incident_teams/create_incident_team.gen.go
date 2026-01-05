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
	Long: `Create a new incident team
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#create-incident-team`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTeamResponse
		var err error

		var body datadogV2.IncidentTeamCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentTeam(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-team")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {

	CreateIncidentTeamCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentTeamCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentTeamCmd)
}
