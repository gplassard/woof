package incident_teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentTeamsCmd = &cobra.Command{
	Use:     "list-incident-teams",
	Aliases: []string{"list"},
	Short:   "Get a list of all incident teams",
	Long: `Get a list of all incident teams
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#list-incident-teams`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTeamsResponse
		var err error

		optionalParams := datadogV2.NewListIncidentTeamsOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentTeams(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-teams")

		fmt.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {

	ListIncidentTeamsCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	ListIncidentTeamsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListIncidentTeamsCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListIncidentTeamsCmd.Flags().String("filter", "", "A search query that filters teams by name.")

	Cmd.AddCommand(ListIncidentTeamsCmd)
}
