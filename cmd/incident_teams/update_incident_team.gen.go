package incident_teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentTeamCmd = &cobra.Command{
	Use:     "update-incident-team [team_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident team",
	Long: `Update an existing incident team
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#update-incident-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTeamResponse
		var err error

		var body datadogV2.IncidentTeamUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentTeam(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-team")

		fmt.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {

	UpdateIncidentTeamCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentTeamCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentTeamCmd)
}
