package okta_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOktaAccountCmd = &cobra.Command{
	Use: "update-okta-account [account_id] [payload]",

	Short: "Update Okta account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OktaAccountResponse
		var err error

		var body datadogV2.OktaAccountUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateOktaAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOktaAccountCmd)
}
