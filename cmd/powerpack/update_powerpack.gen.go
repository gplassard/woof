package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdatePowerpackCmd = &cobra.Command{
	Use:     "update-powerpack [powerpack_id]",
	Aliases: []string{"update"},
	Short:   "Update a powerpack",
	Long: `Update a powerpack
Documentation: https://docs.datadoghq.com/api/latest/powerpack/#update-powerpack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PowerpackResponse
		var err error

		var body datadogV2.Powerpack
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdatePowerpack(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {

	UpdatePowerpackCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdatePowerpackCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdatePowerpackCmd)
}
