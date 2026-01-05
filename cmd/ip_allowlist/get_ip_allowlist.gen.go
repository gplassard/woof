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
	Long: `Get IP Allowlist
Documentation: https://docs.datadoghq.com/api/latest/ip-allowlist/#get-ip-allowlist`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IPAllowlistResponse
		var err error

		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		res, _, err = api.GetIPAllowlist(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-ip-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "ip_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(GetIPAllowlistCmd)
}
