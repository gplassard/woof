package key_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use:   "listcurrentuserapplicationkeys",
	Short: "Get all application keys owned by current user",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.ListCurrentUserApplicationKeys(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listcurrentuserapplicationkeys: %v", err)
		}

		cmdutil.PrintJSON(res, "key_management")
	},
}

func init() {
	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
