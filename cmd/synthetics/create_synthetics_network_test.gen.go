package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSyntheticsNetworkTestCmd = &cobra.Command{
	Use:     "create-synthetics-network-test",
	Aliases: []string{"create-network-test"},
	Short:   "Create synthetics network test",
	Long: `Create synthetics network test
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#create-synthetics-network-test`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsNetworkTestResponse
		var err error

		var body datadogV2.SyntheticsNetworkTestEditRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSyntheticsNetworkTest(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-synthetics-network-test")

		fmt.Println(cmdutil.FormatJSON(res, "network_test"))
	},
}

func init() {

	CreateSyntheticsNetworkTestCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSyntheticsNetworkTestCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSyntheticsNetworkTestCmd)
}
