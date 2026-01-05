package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use: "list-current-user-application-keys",

	Short: "Get all application keys owned by current user",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.ListCurrentUserApplicationKeys(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-current-user-application-keys")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
