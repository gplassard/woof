package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreatePowerpackCmd = &cobra.Command{
	Use:     "create-powerpack [payload]",
	Aliases: []string{"create"},
	Short:   "Create a new powerpack",
	Long: `Create a new powerpack
Documentation: https://docs.datadoghq.com/api/latest/powerpack/#create-powerpack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PowerpackResponse
		var err error

		var body datadogV2.Powerpack
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err = api.CreatePowerpack(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {
	Cmd.AddCommand(CreatePowerpackCmd)
}
