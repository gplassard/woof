package authn_mappings

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use:   "listauthnmappings",
	Short: "List all AuthN Mappings",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err := api.ListAuthNMappings(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listauthnmappings: %v", err)
		}

		cmdutil.PrintJSON(res, "authn_mappings")
	},
}

func init() {
	Cmd.AddCommand(ListAuthNMappingsCmd)
}
