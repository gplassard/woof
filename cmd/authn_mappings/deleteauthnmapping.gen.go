package authn_mappings

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAuthNMappingCmd = &cobra.Command{
	Use:   "deleteauthnmapping [authn_mapping_id]",
	Short: "Delete an AuthN Mapping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		_, err := api.DeleteAuthNMapping(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleteauthnmapping: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteAuthNMappingCmd)
}
