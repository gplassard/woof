package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTeamConnectionsCmd = &cobra.Command{
	Use:     "list-team-connections",
	Aliases: []string{"list-connections"},
	Short:   "List team connections",
	Long: `List team connections
Documentation: https://docs.datadoghq.com/api/latest/teams/#list-team-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamConnectionsResponse
		var err error

		optionalParams := datadogV2.NewListTeamConnectionsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("filter-sources") {
			val, _ := cmd.Flags().GetString("filter-sources")
			optionalParams.WithFilterSources(val)
		}

		if cmd.Flags().Changed("filter-team-ids") {
			val, _ := cmd.Flags().GetString("filter-team-ids")
			optionalParams.WithFilterTeamIds(val)
		}

		if cmd.Flags().Changed("filter-connected-team-ids") {
			val, _ := cmd.Flags().GetString("filter-connected-team-ids")
			optionalParams.WithFilterConnectedTeamIds(val)
		}

		if cmd.Flags().Changed("filter-connection-ids") {
			val, _ := cmd.Flags().GetString("filter-connection-ids")
			optionalParams.WithFilterConnectionIds(val)
		}

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTeamConnections(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-team-connections")

		cmd.Println(cmdutil.FormatJSON(res, "team_connection"))
	},
}

func init() {

	ListTeamConnectionsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListTeamConnectionsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListTeamConnectionsCmd.Flags().String("filter-sources", "", "Filter team connections by external source systems.")

	ListTeamConnectionsCmd.Flags().String("filter-team-ids", "", "Filter team connections by Datadog team IDs.")

	ListTeamConnectionsCmd.Flags().String("filter-connected-team-ids", "", "Filter team connections by connected team IDs from external systems.")

	ListTeamConnectionsCmd.Flags().String("filter-connection-ids", "", "Filter team connections by connection IDs.")

	Cmd.AddCommand(ListTeamConnectionsCmd)
}
