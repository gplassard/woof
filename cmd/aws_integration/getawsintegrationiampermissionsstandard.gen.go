package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAWSIntegrationIAMPermissionsStandardCmd = &cobra.Command{
	Use:   "getawsintegrationiampermissionsstandard",
	Short: "Get AWS integration standard IAM permissions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSIntegrationIAMPermissionsStandard(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getawsintegrationiampermissionsstandard: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_integration")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsStandardCmd)
}
