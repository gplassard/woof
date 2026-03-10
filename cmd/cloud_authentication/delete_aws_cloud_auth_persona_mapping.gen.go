package cloud_authentication

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAWSCloudAuthPersonaMappingCmd = &cobra.Command{
	Use: "delete-aws-cloud-auth-persona-mapping [persona_mapping_id]",

	Short: "Delete an AWS cloud authentication persona mapping",
	Long: `Delete an AWS cloud authentication persona mapping
Documentation: https://docs.datadoghq.com/api/latest/cloud-authentication/#delete-aws-cloud-auth-persona-mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCloudAuthenticationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteAWSCloudAuthPersonaMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-aws-cloud-auth-persona-mapping")

	},
}

func init() {

	Cmd.AddCommand(DeleteAWSCloudAuthPersonaMappingCmd)
}
