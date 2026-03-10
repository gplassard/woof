package cloud_authentication

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSCloudAuthPersonaMappingsCmd = &cobra.Command{
	Use: "list-aws-cloud-auth-persona-mappings",

	Short: "List AWS cloud authentication persona mappings",
	Long: `List AWS cloud authentication persona mappings
Documentation: https://docs.datadoghq.com/api/latest/cloud-authentication/#list-aws-cloud-auth-persona-mappings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSCloudAuthPersonaMappingsResponse
		var err error

		api := datadogV2.NewCloudAuthenticationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAWSCloudAuthPersonaMappings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-cloud-auth-persona-mappings")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_authentication"))
	},
}

func init() {

	Cmd.AddCommand(ListAWSCloudAuthPersonaMappingsCmd)
}
