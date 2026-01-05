package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SendInvitationsCmd = &cobra.Command{
	Use: "send-invitations [payload]",

	Short: "Send invitation emails",
	Long: `Send invitation emails
Documentation: https://docs.datadoghq.com/api/latest/users/#send-invitations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserInvitationsResponse
		var err error

		var body datadogV2.UserInvitationsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err = api.SendInvitations(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to send-invitations")

		cmd.Println(cmdutil.FormatJSON(res, "user_invitations"))
	},
}

func init() {
	Cmd.AddCommand(SendInvitationsCmd)
}
