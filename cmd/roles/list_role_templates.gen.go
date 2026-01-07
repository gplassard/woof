package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRoleTemplatesCmd = &cobra.Command{
	Use:     "list-role-templates",
	Aliases: []string{"list-templates"},
	Short:   "List role templates",
	Long: `List role templates
Documentation: https://docs.datadoghq.com/api/latest/roles/#list-role-templates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleTemplateArray
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRoleTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-role-templates")

		fmt.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	Cmd.AddCommand(ListRoleTemplatesCmd)
}
