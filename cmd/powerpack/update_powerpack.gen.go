package powerpack

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdatePowerpackCmd = &cobra.Command{
	Use:   "update-powerpack [powerpack_id]",
	Short: "Update a powerpack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.UpdatePowerpack(client.NewContext(apiKey, appKey, site), args[0], datadogV2.Powerpack{})
		if err != nil {
			log.Fatalf("failed to update-powerpack: %v", err)
		}

		cmdutil.PrintJSON(res, "powerpack")
	},
}

func init() {
	Cmd.AddCommand(UpdatePowerpackCmd)
}
