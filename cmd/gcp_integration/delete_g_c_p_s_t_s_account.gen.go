package gcp_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteGCPSTSAccountCmd = &cobra.Command{
	Use:   "delete_g_c_p_s_t_s_account [account_id]",
	Short: "Delete an STS enabled GCP Account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteGCPSTSAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_g_c_p_s_t_s_account: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteGCPSTSAccountCmd)
}
