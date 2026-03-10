package service_level_objectives

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetSloStatusCmd = &cobra.Command{
	Use:     "get-slo-status [slo_id] [from_ts] [to_ts]",
	Aliases: []string{"get-status"},
	Short:   "Get SLO status",
	Long: `Get SLO status
Documentation: https://docs.datadoghq.com/api/latest/service-level-objectives/#get-slo-status`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SloStatusResponse
		var err error

		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSloStatus(client.NewContext(apiKey, appKey, site), args[0], func() int64 { i, _ := strconv.ParseInt(args[1], 10, 64); return i }(), func() int64 { i, _ := strconv.ParseInt(args[2], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-slo-status")

		fmt.Println(cmdutil.FormatJSON(res, "slo_status"))
	},
}

func init() {

	Cmd.AddCommand(GetSloStatusCmd)
}
