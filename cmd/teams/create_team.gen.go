package teams

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamCmd = &cobra.Command{
	Use:     "create-team",
	Aliases: []string{"create"},
	Short:   "Create a team",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeam(client.NewContext(apiKey, appKey, site), datadogV2.TeamCreateRequest{})
		cmdutil.HandleError(err, "failed to create-team")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamCmd)
}
