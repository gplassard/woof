package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCurrentUserApplicationKeyCmd = &cobra.Command{
	Use: "delete-current-user-application-key [app_key_id]",

	Short: "Delete an application key owned by current user",
	Long: `Delete an application key owned by current user
Documentation: https://docs.datadoghq.com/api/latest/key-management/#delete-current-user-application-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		_, err = api.DeleteCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-current-user-application-key")

	},
}

func init() {
	Cmd.AddCommand(DeleteCurrentUserApplicationKeyCmd)
}
