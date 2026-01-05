package roles

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRoleTemplatesCmd = &cobra.Command{
	Use:     "list-role-templates",
	Aliases: []string{"list-templates"},
	Short:   "List role templates",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRoleTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-role-templates")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(ListRoleTemplatesCmd)
}
