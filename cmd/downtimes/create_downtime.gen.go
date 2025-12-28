package downtimes

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDowntimeCmd = &cobra.Command{
	Use:   "create-downtime",
	Aliases: []string{ "create", },
	Short: "Schedule a downtime",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.CreateDowntime(client.NewContext(apiKey, appKey, site), datadogV2.DowntimeCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-downtime: %v", err)
		}

		cmdutil.PrintJSON(res, "downtime")
	},
}

func init() {
	Cmd.AddCommand(CreateDowntimeCmd)
}
