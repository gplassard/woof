package service_level_objectives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSLOReportJobCmd = &cobra.Command{
	Use:   "create_s_l_o_report_job",
	Short: "Create a new SLO report",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.CreateSLOReportJob(client.NewContext(apiKey, appKey, site), datadogV2.SloReportCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_s_l_o_report_job: %v", err)
		}

		cmdutil.PrintJSON(res, "service_level_objectives")
	},
}

func init() {
	Cmd.AddCommand(CreateSLOReportJobCmd)
}
