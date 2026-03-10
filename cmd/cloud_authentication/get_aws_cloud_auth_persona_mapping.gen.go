package cloud_authentication

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSCloudAuthPersonaMappingCmd = &cobra.Command{
	Use: "get-aws-cloud-auth-persona-mapping [persona_mapping_id]",

	Short: "Get an AWS cloud authentication persona mapping",
	Long: `Get an AWS cloud authentication persona mapping
Documentation: https://docs.datadoghq.com/api/latest/cloud-authentication/#get-aws-cloud-auth-persona-mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSCloudAuthPersonaMappingResponse
		var err error

		api := datadogV2.NewCloudAuthenticationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAWSCloudAuthPersonaMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-aws-cloud-auth-persona-mapping")

		fmt.Println(cmdutil.FormatJSON(res, "aws_cloud_auth_config"))
	},
}

func init() {

	Cmd.AddCommand(GetAWSCloudAuthPersonaMappingCmd)
}
