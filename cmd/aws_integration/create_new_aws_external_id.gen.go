package aws_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateNewAWSExternalIDCmd = &cobra.Command{
	Use: "create-new-aws-external-id",

	Short: "Generate a new external ID",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateNewAWSExternalID(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to create-new-aws-external-id")

		cmd.Println(cmdutil.FormatJSON(res, "external_id"))
	},
}

func init() {
	Cmd.AddCommand(CreateNewAWSExternalIDCmd)
}
