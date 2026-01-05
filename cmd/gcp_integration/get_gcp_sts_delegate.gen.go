package gcp_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetGCPSTSDelegateCmd = &cobra.Command{
	Use: "get-gcp-sts-delegate",

	Short: "List delegate account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetGCPSTSDelegate(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-gcp-sts-delegate")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_sts_delegate"))
	},
}

func init() {
	Cmd.AddCommand(GetGCPSTSDelegateCmd)
}
