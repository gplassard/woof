package aws_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAWSAccountCmd = &cobra.Command{
	Use: "create-aws-account",

	Short: "Create an AWS integration",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateAWSAccount(client.NewContext(apiKey, appKey, site), datadogV2.AWSAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-aws-account")

		cmd.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {
	Cmd.AddCommand(CreateAWSAccountCmd)
}
