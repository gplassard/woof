package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateNewAWSExternalIDCmd = &cobra.Command{
	Use: "create-new-aws-external-id",

	Short: "Generate a new external ID",
	Long: `Generate a new external ID
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#create-new-aws-external-id`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSNewExternalIDResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateNewAWSExternalID(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to create-new-aws-external-id")

		fmt.Println(cmdutil.FormatJSON(res, "external_id"))
	},
}

func init() {

	Cmd.AddCommand(CreateNewAWSExternalIDCmd)
}
