package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAWSIntegrationIAMPermissionsCmd = &cobra.Command{
	Use:   "get-a-w-s-integration-i-a-m-permissions",
	
	Short: "Get AWS integration IAM permissions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSIntegrationIAMPermissions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-a-w-s-integration-i-a-m-permissions: %v", err)
		}

		cmdutil.PrintJSON(res, "permissions")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsCmd)
}
