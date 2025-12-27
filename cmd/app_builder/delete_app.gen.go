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

var DeleteAppCmd = &cobra.Command{
	Use:   "delete_app [app_id]",
	Short: "Delete App",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.DeleteApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to delete_app: %v", err)
		}

		cmdutil.PrintJSON(res, "app_builder")
	},
}

func init() {
	Cmd.AddCommand(DeleteAppCmd)
}
