package organizations

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOrgConfigsCmd = &cobra.Command{
	Use: "list-org-configs",

	Short: "List Org Configs",
	Long: `List Org Configs
Documentation: https://docs.datadoghq.com/api/latest/organizations/#list-org-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConfigListResponse
		var err error

		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListOrgConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-org-configs")

		fmt.Println(cmdutil.FormatJSON(res, "org_configs"))
	},
}

func init() {

	Cmd.AddCommand(ListOrgConfigsCmd)
}
