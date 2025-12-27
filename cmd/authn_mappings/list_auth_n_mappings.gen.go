package authn_mappings

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use:   "list_auth_n_mappings",
	Short: "List all AuthN Mappings",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err := api.ListAuthNMappings(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_auth_n_mappings: %v", err)
		}

		cmdutil.PrintJSON(res, "authn_mappings")
	},
}

func init() {
	Cmd.AddCommand(ListAuthNMappingsCmd)
}
