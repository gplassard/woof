package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetPowerpackCmd = &cobra.Command{
	Use:     "get-powerpack [powerpack_id]",
	Aliases: []string{"get"},
	Short:   "Get a Powerpack",
	Long: `Get a Powerpack
Documentation: https://docs.datadoghq.com/api/latest/powerpack/#get-powerpack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PowerpackResponse
		var err error

		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetPowerpack(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {

	Cmd.AddCommand(GetPowerpackCmd)
}
