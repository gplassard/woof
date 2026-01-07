package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use: "list-current-user-application-keys",

	Short: "Get all application keys owned by current user",
	Long: `Get all application keys owned by current user
Documentation: https://docs.datadoghq.com/api/latest/key-management/#list-current-user-application-keys`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListApplicationKeysResponse
		var err error

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCurrentUserApplicationKeys(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-current-user-application-keys")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
