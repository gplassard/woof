package a_w_s_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsCmd = &cobra.Command{
	Use: "get-aws-integration-iam-permissions",

	Short: "Get AWS integration IAM permissions",
	Long: `Get AWS integration IAM permissions
Documentation: https://docs.datadoghq.com/api/latest/a-w-s-integration/#get-aws-integration-iam-permissions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSIntegrationIamPermissionsResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAWSIntegrationIAMPermissions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aws-integration-iam-permissions")

		fmt.Println(cmdutil.FormatJSON(res, "a_w_s_integration_i_a_m_permission"))
	},
}

func init() {

	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsCmd)
}
