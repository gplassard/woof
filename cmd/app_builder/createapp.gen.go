package app_builder

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAppCmd = &cobra.Command{
	Use:   "createapp",
	Short: "Create App",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.CreateApp(client.NewContext(apiKey, appKey, site), datadogV2.CreateAppRequest{})
		if err != nil {
			log.Fatalf("failed to createapp: %v", err)
		}

		cmdutil.PrintJSON(res, "app_builder")
	},
}

func init() {
	Cmd.AddCommand(CreateAppCmd)
}
