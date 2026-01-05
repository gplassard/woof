package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAWSAccountCmd = &cobra.Command{
	Use: "update-aws-account [aws_account_config_id] [payload]",

	Short: "Update an AWS integration",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.AWSAccountUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateAWSAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-aws-account")

		cmd.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAWSAccountCmd)
}
