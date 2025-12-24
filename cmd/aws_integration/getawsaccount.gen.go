package aws_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAWSAccountCmd = &cobra.Command{
	Use:   "getawsaccount [aws_account_config_id]",
	Short: "Get an AWS integration by config ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetAWSAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getawsaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_integration")
	},
}

func init() {
	Cmd.AddCommand(GetAWSAccountCmd)
}
