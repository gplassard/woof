package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddMemberTeamCmd = &cobra.Command{
	Use:     "add-member-team [super_team_id]",
	Aliases: []string{"add-member"},
	Short:   "Add a member team",
	Long: `Add a member team
Documentation: https://docs.datadoghq.com/api/latest/teams/#add-member-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.AddMemberTeamRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.AddMemberTeam(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-member-team")

	},
}

func init() {

	AddMemberTeamCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AddMemberTeamCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AddMemberTeamCmd)
}
