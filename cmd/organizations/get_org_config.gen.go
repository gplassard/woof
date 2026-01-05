package organizations

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOrgConfigCmd = &cobra.Command{
	Use: "get-org-config [org_config_name]",

	Short: "Get a specific Org Config value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err := api.GetOrgConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-org-config")

		cmd.Println(cmdutil.FormatJSON(res, "org_configs"))
	},
}

func init() {
	Cmd.AddCommand(GetOrgConfigCmd)
}
