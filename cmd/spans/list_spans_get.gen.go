package spans

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSpansGetCmd = &cobra.Command{
	Use:   "list-spans-get",
	Aliases: []string{ "list-get", },
	Short: "Get a list of spans",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.ListSpansGet(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-spans-get: %v", err)
		}

		cmdutil.PrintJSON(res, "spans")
	},
}

func init() {
	Cmd.AddCommand(ListSpansGetCmd)
}
