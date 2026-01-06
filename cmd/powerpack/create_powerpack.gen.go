package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreatePowerpackCmd = &cobra.Command{
	Use:     "create-powerpack",
	Aliases: []string{"create"},
	Short:   "Create a new powerpack",
	Long: `Create a new powerpack
Documentation: https://docs.datadoghq.com/api/latest/powerpack/#create-powerpack`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PowerpackResponse
		var err error

		var body datadogV2.Powerpack
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreatePowerpack(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {

	CreatePowerpackCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreatePowerpackCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreatePowerpackCmd)
}
