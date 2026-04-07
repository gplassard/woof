package gcp_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var MakeGCPSTSDelegateCmd = &cobra.Command{
	Use: "make-gcp-sts-delegate",

	Short: "Create a Datadog GCP principal",
	Long: `Create a Datadog GCP principal
Documentation: https://docs.datadoghq.com/api/latest/gcp-integration/#make-gcp-sts-delegate`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSDelegateAccountResponse
		var err error

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.MakeGCPSTSDelegate(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to make-gcp-sts-delegate")

		fmt.Println(cmdutil.FormatJSON(res, "gcp_sts_delegate"))
	},
}

func init() {

	Cmd.AddCommand(MakeGCPSTSDelegateCmd)
}
