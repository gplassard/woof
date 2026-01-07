package gcp_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetGCPSTSDelegateCmd = &cobra.Command{
	Use: "get-gcp-sts-delegate",

	Short: "List delegate account",
	Long: `List delegate account
Documentation: https://docs.datadoghq.com/api/latest/gcp-integration/#get-gcp-sts-delegate`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSDelegateAccountResponse
		var err error

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetGCPSTSDelegate(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-gcp-sts-delegate")

		fmt.Println(cmdutil.FormatJSON(res, "gcp_sts_delegate"))
	},
}

func init() {

	Cmd.AddCommand(GetGCPSTSDelegateCmd)
}
