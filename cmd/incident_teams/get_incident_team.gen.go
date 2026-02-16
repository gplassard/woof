package incident_teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentTeamCmd = &cobra.Command{
	Use:     "get-incident-team [team_id]",
	Aliases: []string{"get"},
	Short:   "Get details of an incident team",
	Long: `Get details of an incident team
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#get-incident-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTeamResponse
		var err error

		optionalParams := datadogV2.NewGetIncidentTeamOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetIncidentTeam(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-incident-team")

		fmt.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {

	GetIncidentTeamCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	Cmd.AddCommand(GetIncidentTeamCmd)
}
