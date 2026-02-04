package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsStandardCmd = &cobra.Command{
	Use:     "get-aws-integration-iam-permissions-standard",
	Aliases: []string{"get-iam-permissions-standard"},
	Short:   "Get AWS integration standard IAM permissions",
	Long: `Get AWS integration standard IAM permissions
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#get-aws-integration-iam-permissions-standard`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSIntegrationIamPermissionsResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAWSIntegrationIAMPermissionsStandard(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aws-integration-iam-permissions-standard")

		fmt.Println(cmdutil.FormatJSON(res, "a_w_s_integration_i_a_m_permissions_standard"))
	},
}

func init() {

	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsStandardCmd)
}
