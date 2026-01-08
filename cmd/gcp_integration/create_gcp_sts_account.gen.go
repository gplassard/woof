package gcp_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateGCPSTSAccountCmd = &cobra.Command{
	Use: "create-gcp-sts-account",

	Short: "Create a new entry for your service account",
	Long: `Create a new entry for your service account
Documentation: https://docs.datadoghq.com/api/latest/g-c-p-integration/#create-gcp-sts-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSServiceAccountResponse
		var err error

		var body datadogV2.GCPSTSServiceAccountCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateGCPSTSAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-gcp-sts-account")

		fmt.Println(cmdutil.FormatJSON(res, "g_c_p_s_t_s_account"))
	},
}

func init() {

	CreateGCPSTSAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateGCPSTSAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateGCPSTSAccountCmd)
}
