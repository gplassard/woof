package ip_allowlist

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIPAllowlistCmd = &cobra.Command{
	Use:     "update-ip-allowlist [payload]",
	Aliases: []string{"update"},
	Short:   "Update IP Allowlist",
	Long: `Update IP Allowlist
Documentation: https://docs.datadoghq.com/api/latest/ip-allowlist/#update-ip-allowlist`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IPAllowlistResponse
		var err error

		var body datadogV2.IPAllowlistUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		res, _, err = api.UpdateIPAllowlist(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-ip-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "ip_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIPAllowlistCmd)
}
