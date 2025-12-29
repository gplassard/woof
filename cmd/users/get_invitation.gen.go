package users

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetInvitationCmd = &cobra.Command{
	Use:   "get-invitation [user_invitation_uuid]",
	
	Short: "Get a user invitation",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.GetInvitation(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-invitation")

		cmd.Println(cmdutil.FormatJSON(res, "user_invitations"))
	},
}

func init() {
	Cmd.AddCommand(GetInvitationCmd)
}
