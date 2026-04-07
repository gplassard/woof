package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamNotificationRuleCmd = &cobra.Command{
	Use:     "create-team-notification-rule [team_id]",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create team notification rule",
	Long: `Create team notification rule
Documentation: https://docs.datadoghq.com/api/latest/teams/#create-team-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamNotificationRuleResponse
		var err error

		var body datadogV2.TeamNotificationRuleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-team-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {

	CreateTeamNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTeamNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTeamNotificationRuleCmd)
}
