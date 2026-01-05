package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsResourceCollectionCmd = &cobra.Command{
	Use:     "get-aws-integration-iam-permissions-resource-collection",
	Aliases: []string{"get-iam-permissions-resource-collection"},
	Short:   "Get resource collection IAM permissions",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSIntegrationIAMPermissionsResourceCollection(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aws-integration-iam-permissions-resource-collection")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsResourceCollectionCmd)
}
