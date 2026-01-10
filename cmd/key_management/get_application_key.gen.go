package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetApplicationKeyCmd = &cobra.Command{
	Use: "get-application-key [app_key_id]",

	Short: "Get an application key",
	Long: `Get an application key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#get-application-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationKeyResponse
		var err error

		optionalParams := datadogV2.NewGetApplicationKeyOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetApplicationKey(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-application-key")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	GetApplicationKeyCmd.Flags().String("include", "", "Resource path for related resources to include in the response. Only 'owned_by' is supported.")

	Cmd.AddCommand(GetApplicationKeyCmd)
}
