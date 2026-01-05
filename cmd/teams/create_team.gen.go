package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamCmd = &cobra.Command{
	Use:     "create-team",
	Aliases: []string{"create"},
	Short:   "Create a team",
	Long: `Create a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#create-team`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamResponse
		var err error

		var body datadogV2.TeamCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.CreateTeam(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-team")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {

	CreateTeamCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTeamCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTeamCmd)
}
