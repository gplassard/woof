package gcp_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteGCPSTSAccountCmd = &cobra.Command{
	Use: "delete-gcp-sts-account [account_id]",

	Short: "Delete an STS enabled GCP Account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		_, err = api.DeleteGCPSTSAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-gcp-sts-account")

	},
}

func init() {
	Cmd.AddCommand(DeleteGCPSTSAccountCmd)
}
