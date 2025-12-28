package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateApplicationKeyCmd = &cobra.Command{
	Use:   "update_application_key [app_key_id]",
	Short: "Edit an application key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateApplicationKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationKeyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_application_key: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationKeyCmd)
}
