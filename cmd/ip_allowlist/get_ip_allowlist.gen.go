package ip_allowlist

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIPAllowlistCmd = &cobra.Command{
	Use:     "get-ip-allowlist",
	Aliases: []string{"get"},
	Short:   "Get IP Allowlist",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		res, _, err := api.GetIPAllowlist(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-ip-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "ip_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(GetIPAllowlistCmd)
}
