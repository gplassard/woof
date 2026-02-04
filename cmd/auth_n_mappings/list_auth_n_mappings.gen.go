package auth_n_mappings

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use:     "list-auth-n-mappings",
	Aliases: []string{"list"},
	Short:   "List all AuthN Mappings",
	Long: `List all AuthN Mappings
Documentation: https://docs.datadoghq.com/api/latest/auth-n-mappings/#list-auth-n-mappings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingsResponse
		var err error

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAuthNMappings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-auth-n-mappings")

		fmt.Println(cmdutil.FormatJSON(res, "auth_n_mapping"))
	},
}

func init() {

	Cmd.AddCommand(ListAuthNMappingsCmd)
}
