package ip_allowlist

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIPAllowlistCmd = &cobra.Command{
	Use:     "update-ip-allowlist",
	Aliases: []string{"update"},
	Short:   "Update IP Allowlist",
	Long: `Update IP Allowlist
Documentation: https://docs.datadoghq.com/api/latest/ip-allowlist/#update-ip-allowlist`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IPAllowlistResponse
		var err error

		var body datadogV2.IPAllowlistUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIPAllowlist(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-ip-allowlist")

		fmt.Println(cmdutil.FormatJSON(res, "ip_allowlist"))
	},
}

func init() {

	UpdateIPAllowlistCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIPAllowlistCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIPAllowlistCmd)
}
