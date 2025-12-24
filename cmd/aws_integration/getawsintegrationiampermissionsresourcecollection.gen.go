package aws_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAWSIntegrationIAMPermissionsResourceCollectionCmd = &cobra.Command{
	Use:   "getawsintegrationiampermissionsresourcecollection",
	Short: "Get resource collection IAM permissions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSIntegrationIAMPermissionsResourceCollection(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getawsintegrationiampermissionsresourcecollection: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_integration")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsResourceCollectionCmd)
}
