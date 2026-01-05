package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use: "update-current-user-application-key [app_key_id]",

	Short: "Edit an application key owned by current user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationKeyUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-current-user-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCurrentUserApplicationKeyCmd)
}
