package powerpack

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListPowerpacksCmd = &cobra.Command{
	Use:   "listpowerpacks",
	Short: "Get all powerpacks",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.ListPowerpacks(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listpowerpacks: %v", err)
		}

		cmdutil.PrintJSON(res, "powerpack")
	},
}

func init() {
	Cmd.AddCommand(ListPowerpacksCmd)
}
