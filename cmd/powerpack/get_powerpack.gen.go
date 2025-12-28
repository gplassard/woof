package powerpack

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetPowerpackCmd = &cobra.Command{
	Use:   "get-powerpack [powerpack_id]",
	Aliases: []string{ "get", },
	Short: "Get a Powerpack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.GetPowerpack(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-powerpack")

		cmdutil.PrintJSON(res, "powerpack")
	},
}

func init() {
	Cmd.AddCommand(GetPowerpackCmd)
}
