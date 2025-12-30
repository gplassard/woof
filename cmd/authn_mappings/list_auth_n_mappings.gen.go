package authn_mappings

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use: "list-auth-n-mappings",

	Short: "List all AuthN Mappings",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err := api.ListAuthNMappings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-auth-n-mappings")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {
	Cmd.AddCommand(ListAuthNMappingsCmd)
}
