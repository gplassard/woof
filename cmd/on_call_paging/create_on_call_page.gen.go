package on_call_paging

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOnCallPageCmd = &cobra.Command{
	Use: "create-on-call-page",

	Short: "Create On-Call Page",
	Long: `Create On-Call Page
Documentation: https://docs.datadoghq.com/api/latest/on-call-paging/#create-on-call-page`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreatePageResponse
		var err error

		var body datadogV2.CreatePageRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOnCallPage(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-on-call-page")

		fmt.Println(cmdutil.FormatJSON(res, "on_call_page"))
	},
}

func init() {

	CreateOnCallPageCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOnCallPageCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOnCallPageCmd)
}
