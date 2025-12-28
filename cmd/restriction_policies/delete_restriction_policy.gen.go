package restriction_policies

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRestrictionPolicyCmd = &cobra.Command{
	Use:   "delete-restriction-policy [resource_id]",
	
	Short: "Delete a restriction policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		_, err := api.DeleteRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-restriction-policy: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionPolicyCmd)
}
