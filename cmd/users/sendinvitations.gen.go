package users

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SendInvitationsCmd = &cobra.Command{
	Use:   "sendinvitations",
	Short: "Send invitation emails",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.SendInvitations(client.NewContext(apiKey, appKey, site), datadogV2.UserInvitationsRequest{})
		if err != nil {
			log.Fatalf("failed to sendinvitations: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}

func init() {
	Cmd.AddCommand(SendInvitationsCmd)
}
