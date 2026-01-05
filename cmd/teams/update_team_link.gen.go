package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamLinkCmd = &cobra.Command{
	Use:     "update-team-link [team_id] [link_id]",
	Aliases: []string{"update-link"},
	Short:   "Update a team link",
	Long: `Update a team link
Documentation: https://docs.datadoghq.com/api/latest/teams/#update-team-link`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamLinkResponse
		var err error

		var body datadogV2.TeamLinkCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.UpdateTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_links"))
	},
}

func init() {

	UpdateTeamLinkCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTeamLinkCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTeamLinkCmd)
}
