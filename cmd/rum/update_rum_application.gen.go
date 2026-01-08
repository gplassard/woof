package rum

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRUMApplicationCmd = &cobra.Command{
	Use:     "update-rum-application [id]",
	Aliases: []string{"update-application"},
	Short:   "Update a RUM application",
	Long: `Update a RUM application
Documentation: https://docs.datadoghq.com/api/latest/r-u-m/#update-rum-application`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMApplicationResponse
		var err error

		var body datadogV2.RUMApplicationUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRUMApplication(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-rum-application")

		fmt.Println(cmdutil.FormatJSON(res, "r_u_m_application"))
	},
}

func init() {

	UpdateRUMApplicationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRUMApplicationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRUMApplicationCmd)
}
