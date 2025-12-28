package app_builder

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var GetAppCmd = &cobra.Command{
	Use:   "get-app [app_id]",
	
	Short: "Get App",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.GetApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to get-app: %v", err)
		}

		cmdutil.PrintJSON(res, "appDefinitions")
	},
}

func init() {
	Cmd.AddCommand(GetAppCmd)
}
