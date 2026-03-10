package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ActivateContentPackCmd = &cobra.Command{
	Use: "activate-content-pack [content_pack_id]",

	Short: "Activate content pack",
	Long: `Activate content pack
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#activate-content-pack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ActivateContentPack(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to activate-content-pack")

	},
}

func init() {

	Cmd.AddCommand(ActivateContentPackCmd)
}
