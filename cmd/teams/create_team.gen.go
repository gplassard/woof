package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTeamCmd = &cobra.Command{
	Use:     "create-team [payload]",
	Aliases: []string{"create"},
	Short:   "Create a team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamResponse
		var err error

		var body datadogV2.TeamCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.CreateTeam(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-team")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamCmd)
}
