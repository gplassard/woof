package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRoleCmd = &cobra.Command{
	Use:     "create-role",
	Aliases: []string{"create"},
	Short:   "Create role",
	Long: `Create role
Documentation: https://docs.datadoghq.com/api/latest/roles/#create-role`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleCreateResponse
		var err error

		var body datadogV2.RoleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRole(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-role")

		fmt.Println(cmdutil.FormatJSON(res, "role"))
	},
}

func init() {

	CreateRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRoleCmd)
}
