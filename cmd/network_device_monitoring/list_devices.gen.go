package network_device_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDevicesCmd = &cobra.Command{
	Use: "list-devices",

	Short: "Get the list of devices",
	Long: `Get the list of devices
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#list-devices`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListDevicesResponse
		var err error

		optionalParams := datadogV2.NewListDevicesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter-tag") {
			val, _ := cmd.Flags().GetString("filter-tag")
			optionalParams.WithFilterTag(val)
		}

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDevices(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-devices")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {

	ListDevicesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListDevicesCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListDevicesCmd.Flags().String("sort", "", "The field to sort the devices by.")

	ListDevicesCmd.Flags().String("filter-tag", "", "Filter devices by tag.")

	Cmd.AddCommand(ListDevicesCmd)
}
