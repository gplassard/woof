package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSyntheticsSuiteCmd = &cobra.Command{
	Use:     "get-synthetics-suite [public_id]",
	Aliases: []string{"get-suite"},
	Short:   "GetSyntheticsSuite",
	Long: `GetSyntheticsSuite
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-synthetics-suite`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsSuiteResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSyntheticsSuite(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-synthetics-suite")

		fmt.Println(cmdutil.FormatJSON(res, "suites"))
	},
}

func init() {

	Cmd.AddCommand(GetSyntheticsSuiteCmd)
}
