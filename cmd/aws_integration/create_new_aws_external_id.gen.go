package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateNewAWSExternalIDCmd = &cobra.Command{
	Use:   "create-new-aws-external-id",
	
	Short: "Generate a new external ID",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateNewAWSExternalID(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to create-new-aws-external-id: %v", err)
		}

		cmdutil.PrintJSON(res, "external_id")
	},
}

func init() {
	Cmd.AddCommand(CreateNewAWSExternalIDCmd)
}
