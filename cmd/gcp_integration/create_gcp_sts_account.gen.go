package gcp_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateGCPSTSAccountCmd = &cobra.Command{
	Use: "create-gcp-sts-account [payload]",

	Short: "Create a new entry for your service account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.GCPSTSServiceAccountCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateGCPSTSAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-gcp-sts-account")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_service_account"))
	},
}

func init() {
	Cmd.AddCommand(CreateGCPSTSAccountCmd)
}
