package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamCmd = &cobra.Command{
	Use:     "update-team [team_id]",
	Aliases: []string{"update"},
	Short:   "Update a team",
	Long: `Update a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#update-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamResponse
		var err error

		var body datadogV2.TeamUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateTeam(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-team")

		fmt.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {

	UpdateTeamCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTeamCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTeamCmd)
}
