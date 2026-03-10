package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetContentPacksStatesCmd = &cobra.Command{
	Use: "get-content-packs-states",

	Short: "Get content pack states",
	Long: `Get content pack states
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-content-packs-states`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringContentPackStatesResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetContentPacksStates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-content-packs-states")

		fmt.Println(cmdutil.FormatJSON(res, "content_pack_state"))
	},
}

func init() {

	Cmd.AddCommand(GetContentPacksStatesCmd)
}
