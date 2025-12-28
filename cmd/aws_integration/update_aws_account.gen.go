package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAWSAccountCmd = &cobra.Command{
	Use:   "update-aws-account [aws_account_config_id]",
	
	Short: "Update an AWS integration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateAWSAccount(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AWSAccountUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-aws-account: %v", err)
		}

		cmdutil.PrintJSON(res, "account")
	},
}

func init() {
	Cmd.AddCommand(UpdateAWSAccountCmd)
}
