package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRolesCmd = &cobra.Command{
	Use:     "list-roles",
	Aliases: []string{"list"},
	Short:   "List roles",
	Long: `List roles
Documentation: https://docs.datadoghq.com/api/latest/roles/#list-roles`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RolesResponse
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRoles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-roles")

		fmt.Println(cmdutil.FormatJSON(res, "role"))
	},
}

func init() {

	Cmd.AddCommand(ListRolesCmd)
}
