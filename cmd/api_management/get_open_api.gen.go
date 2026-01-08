package api_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
	"io"
)

var GetOpenAPICmd = &cobra.Command{
	Use: "get-open-api [id]",

	Short: "Get an API",
	Long: `Get an API
Documentation: https://docs.datadoghq.com/api/latest/a-p-i-management/#get-open-api`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res io.Reader
		var err error

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-open-api")

		fmt.Println(cmdutil.FormatJSON(res, "open_a_p_i"))
	},
}

func init() {

	Cmd.AddCommand(GetOpenAPICmd)
}
