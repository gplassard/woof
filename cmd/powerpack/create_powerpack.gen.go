package powerpack

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreatePowerpackCmd = &cobra.Command{
	Use:   "create-powerpack",
	Aliases: []string{ "create", },
	Short: "Create a new powerpack",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.CreatePowerpack(client.NewContext(apiKey, appKey, site), datadogV2.Powerpack{})
		cmdutil.HandleError(err, "failed to create-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {
	Cmd.AddCommand(CreatePowerpackCmd)
}
