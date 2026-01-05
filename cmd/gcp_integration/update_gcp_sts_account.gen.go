package gcp_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateGCPSTSAccountCmd = &cobra.Command{
	Use: "update-gcp-sts-account [account_id]",

	Short: "Update STS Service Account",
	Long: `Update STS Service Account
Documentation: https://docs.datadoghq.com/api/latest/gcp-integration/#update-gcp-sts-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSServiceAccountResponse
		var err error

		var body datadogV2.GCPSTSServiceAccountUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateGCPSTSAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-gcp-sts-account")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_service_account"))
	},
}

func init() {

	UpdateGCPSTSAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateGCPSTSAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateGCPSTSAccountCmd)
}
