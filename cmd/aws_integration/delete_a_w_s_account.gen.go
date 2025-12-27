package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAWSAccountCmd = &cobra.Command{
	Use:   "delete_a_w_s_account [aws_account_config_id]",
	Short: "Delete an AWS integration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteAWSAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_a_w_s_account: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSAccountCmd)
}
