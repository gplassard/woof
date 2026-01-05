package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateConfluentAccountCmd = &cobra.Command{
	Use: "update-confluent-account [account_id] [payload]",

	Short: "Update Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentAccountResponse
		var err error

		var body datadogV2.ConfluentAccountUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err = api.UpdateConfluentAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-confluent-account")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-accounts"))
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentAccountCmd)
}
