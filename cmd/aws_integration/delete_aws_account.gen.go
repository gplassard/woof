package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAWSAccountCmd = &cobra.Command{
	Use:   "delete-aws-account [aws_account_config_id]",
	
	Short: "Delete an AWS integration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteAWSAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-aws-account")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSAccountCmd)
}
