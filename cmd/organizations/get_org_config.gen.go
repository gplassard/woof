package organizations

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOrgConfigCmd = &cobra.Command{
	Use: "get-org-config [org_config_name]",

	Short: "Get a specific Org Config value",
	Long: `Get a specific Org Config value
Documentation: https://docs.datadoghq.com/api/latest/organizations/#get-org-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConfigGetResponse
		var err error

		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetOrgConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-org-config")

		fmt.Println(cmdutil.FormatJSON(res, "org_configs"))
	},
}

func init() {

	Cmd.AddCommand(GetOrgConfigCmd)
}
