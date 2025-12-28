package spans

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSpansCmd = &cobra.Command{
	Use:   "list-spans",
	Aliases: []string{ "list", },
	Short: "Search spans",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.ListSpans(client.NewContext(apiKey, appKey, site), datadogV2.SpansListRequest{})
		if err != nil {
			log.Fatalf("failed to list-spans: %v", err)
		}

		cmdutil.PrintJSON(res, "spans")
	},
}

func init() {
	Cmd.AddCommand(ListSpansCmd)
}
