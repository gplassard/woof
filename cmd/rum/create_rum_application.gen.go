package rum

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRUMApplicationCmd = &cobra.Command{
	Use:     "create-rum-application",
	Aliases: []string{"create-application"},
	Short:   "Create a new RUM application",
	Long: `Create a new RUM application
Documentation: https://docs.datadoghq.com/api/latest/r-u-m/#create-rum-application`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMApplicationResponse
		var err error

		var body datadogV2.RUMApplicationCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRUMApplication(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-rum-application")

		fmt.Println(cmdutil.FormatJSON(res, "r_u_m_application"))
	},
}

func init() {

	CreateRUMApplicationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRUMApplicationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRUMApplicationCmd)
}
