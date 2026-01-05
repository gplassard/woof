package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateConfluentAccountCmd = &cobra.Command{
	Use: "create-confluent-account",

	Short: "Add Confluent account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.CreateConfluentAccount(client.NewContext(apiKey, appKey, site), datadogV2.ConfluentAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-confluent-account")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentAccountCmd)
}
