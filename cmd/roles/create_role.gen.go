package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateRoleCmd = &cobra.Command{
	Use:     "create-role [payload]",
	Aliases: []string{"create"},
	Short:   "Create role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleCreateResponse
		var err error

		var body datadogV2.RoleCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.CreateRole(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(CreateRoleCmd)
}
