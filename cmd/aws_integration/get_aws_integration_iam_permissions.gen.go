package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAWSIntegrationIAMPermissionsCmd = &cobra.Command{
	Use:   "get-aws-integration-iam-permissions",
	Aliases: []string{ "get-iam-permissions", },
	Short: "Get AWS integration IAM permissions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSIntegrationIAMPermissions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aws-integration-iam-permissions")

		cmdutil.PrintJSON(res, "permissions")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsCmd)
}
