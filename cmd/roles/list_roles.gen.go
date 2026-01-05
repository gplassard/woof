package roles

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRolesCmd = &cobra.Command{
	Use:     "list-roles",
	Aliases: []string{"list"},
	Short:   "List roles",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRoles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-roles")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(ListRolesCmd)
}
