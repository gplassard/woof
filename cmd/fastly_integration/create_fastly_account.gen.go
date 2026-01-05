package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateFastlyAccountCmd = &cobra.Command{
	Use: "create-fastly-account [payload]",

	Short: "Add Fastly account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyAccountResponse
		var err error

		var body datadogV2.FastlyAccountCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err = api.CreateFastlyAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-fastly-account")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyAccountCmd)
}
