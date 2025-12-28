package powerpack

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListPowerpacksCmd = &cobra.Command{
	Use:   "list-powerpacks",
	Aliases: []string{ "list-s", },
	Short: "Get all powerpacks",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.ListPowerpacks(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-powerpacks: %v", err)
		}

		cmdutil.PrintJSON(res, "powerpack")
	},
}

func init() {
	Cmd.AddCommand(ListPowerpacksCmd)
}
