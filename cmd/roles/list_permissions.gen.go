package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListPermissionsCmd = &cobra.Command{
	Use: "list-permissions",

	Short: "List permissions",
	Long: `List permissions
Documentation: https://docs.datadoghq.com/api/latest/roles/#list-permissions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PermissionsResponse
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListPermissions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-permissions")

		fmt.Println(cmdutil.FormatJSON(res, "permission"))
	},
}

func init() {

	Cmd.AddCommand(ListPermissionsCmd)
}
