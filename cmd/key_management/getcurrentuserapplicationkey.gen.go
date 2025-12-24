package key_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "getcurrentuserapplicationkey [app_key_id]",
	Short: "Get one application key owned by current user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.GetCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getcurrentuserapplicationkey: %v", err)
		}

		cmdutil.PrintJSON(res, "key_management")
	},
}

func init() {
	Cmd.AddCommand(GetCurrentUserApplicationKeyCmd)
}
