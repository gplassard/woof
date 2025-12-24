package aws_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAWSAccountCmd = &cobra.Command{
	Use:   "createawsaccount",
	Short: "Create an AWS integration",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateAWSAccount(client.NewContext(apiKey, appKey, site), datadogV2.AWSAccountCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createawsaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateAWSAccountCmd)
}
