package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetApiMultistepSubtestsCmd = &cobra.Command{
	Use: "get-api-multistep-subtests [public_id]",

	Short: "Get available subtests for a multistep test",
	Long: `Get available subtests for a multistep test
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-api-multistep-subtests`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsApiMultistepSubtestsResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetApiMultistepSubtests(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-api-multistep-subtests")

		fmt.Println(cmdutil.FormatJSON(res, "subtest"))
	},
}

func init() {

	Cmd.AddCommand(GetApiMultistepSubtestsCmd)
}
