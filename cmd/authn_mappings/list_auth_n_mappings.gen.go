package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use: "list-auth-n-mappings",

	Short: "List all AuthN Mappings",
	Long: `List all AuthN Mappings
Documentation: https://docs.datadoghq.com/api/latest/authn-mappings/#list-auth-n-mappings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingsResponse
		var err error

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAuthNMappings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-auth-n-mappings")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {

	Cmd.AddCommand(ListAuthNMappingsCmd)
}
