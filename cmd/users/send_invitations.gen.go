package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SendInvitationsCmd = &cobra.Command{
	Use: "send-invitations",

	Short: "Send invitation emails",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.SendInvitations(client.NewContext(apiKey, appKey, site), datadogV2.UserInvitationsRequest{})
		cmdutil.HandleError(err, "failed to send-invitations")

		cmd.Println(cmdutil.FormatJSON(res, "user_invitations"))
	},
}

func init() {
	Cmd.AddCommand(SendInvitationsCmd)
}
