package restriction_policies

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRestrictionPolicyCmd = &cobra.Command{
	Use:   "getrestrictionpolicy [resource_id]",
	Short: "Get a restriction policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		res, _, err := api.GetRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getrestrictionpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "restriction_policies")
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionPolicyCmd)
}
