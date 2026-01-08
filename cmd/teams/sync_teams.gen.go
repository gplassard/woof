package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SyncTeamsCmd = &cobra.Command{
	Use:     "sync-teams",
	Aliases: []string{"sync"},
	Short:   "Link Teams with GitHub Teams",
	Long: `Link Teams with GitHub Teams
Documentation: https://docs.datadoghq.com/api/latest/teams/#sync-teams`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.TeamSyncRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.SyncTeams(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to sync-teams")

	},
}

func init() {

	SyncTeamsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SyncTeamsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SyncTeamsCmd)
}
