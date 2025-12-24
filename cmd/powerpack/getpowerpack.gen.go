package powerpack

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetPowerpackCmd = &cobra.Command{
	Use:   "getpowerpack [powerpack_id]",
	Short: "Get a Powerpack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.GetPowerpack(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getpowerpack: %v", err)
		}

		cmdutil.PrintJSON(res, "powerpack")
	},
}

func init() {
	Cmd.AddCommand(GetPowerpackCmd)
}
