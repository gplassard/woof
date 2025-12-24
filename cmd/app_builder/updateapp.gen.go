package app_builder

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAppCmd = &cobra.Command{
	Use:   "updateapp [app_id]",
	Short: "Update App",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		res, _, err := api.UpdateApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.UpdateAppRequest{})
		if err != nil {
			log.Fatalf("failed to updateapp: %v", err)
		}

		cmdutil.PrintJSON(res, "app_builder")
	},
}

func init() {
	Cmd.AddCommand(UpdateAppCmd)
}
