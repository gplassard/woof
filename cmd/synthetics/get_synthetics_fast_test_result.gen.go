package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSyntheticsFastTestResultCmd = &cobra.Command{
	Use:     "get-synthetics-fast-test-result [id]",
	Aliases: []string{"get-fast-test-result"},
	Short:   "Get synthetics fast test result",
	Long: `Get synthetics fast test result
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-synthetics-fast-test-result`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsFastTestResult
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSyntheticsFastTestResult(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-synthetics-fast-test-result")

		fmt.Println(cmdutil.FormatJSON(res, "result"))
	},
}

func init() {

	Cmd.AddCommand(GetSyntheticsFastTestResultCmd)
}
