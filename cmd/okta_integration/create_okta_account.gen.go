package okta_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOktaAccountCmd = &cobra.Command{
	Use: "create-okta-account [payload]",

	Short: "Add Okta account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.OktaAccountRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateOktaAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateOktaAccountCmd)
}
