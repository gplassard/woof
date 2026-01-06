package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSNamespacesCmd = &cobra.Command{
	Use: "list-aws-namespaces",

	Short: "List available namespaces",
	Long: `List available namespaces
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#list-aws-namespaces`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSNamespacesResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAWSNamespaces(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-namespaces")

		cmd.Println(cmdutil.FormatJSON(res, "namespaces"))
	},
}

func init() {

	Cmd.AddCommand(ListAWSNamespacesCmd)
}
