package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DisableUserCmd = &cobra.Command{
	Use:     "disable-user [user_id]",
	Aliases: []string{"disable"},
	Short:   "Disable a user",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		_, err := api.DisableUser(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to disable-user")

	},
}

func init() {
	Cmd.AddCommand(DisableUserCmd)
}
