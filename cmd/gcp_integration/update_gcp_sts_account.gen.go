package gcp_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateGCPSTSAccountCmd = &cobra.Command{
	Use: "update-gcp-sts-account [account_id] [payload]",

	Short: "Update STS Service Account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSServiceAccountResponse
		var err error

		var body datadogV2.GCPSTSServiceAccountUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateGCPSTSAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-gcp-sts-account")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_service_account"))
	},
}

func init() {
	Cmd.AddCommand(UpdateGCPSTSAccountCmd)
}
