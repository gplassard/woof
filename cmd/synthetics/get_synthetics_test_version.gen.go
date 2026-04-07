package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetSyntheticsTestVersionCmd = &cobra.Command{
	Use:     "get-synthetics-test-version [public_id] [version_number]",
	Aliases: []string{"get-test-version"},
	Short:   "Get a specific version of a test",
	Long: `Get a specific version of a test
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-synthetics-test-version`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsTestVersionResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSyntheticsTestVersion(client.NewContext(apiKey, appKey, site), args[0], func() int64 { i, _ := strconv.ParseInt(args[1], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-synthetics-test-version")

		fmt.Println(cmdutil.FormatJSON(res, "version"))
	},
}

func init() {

	Cmd.AddCommand(GetSyntheticsTestVersionCmd)
}
