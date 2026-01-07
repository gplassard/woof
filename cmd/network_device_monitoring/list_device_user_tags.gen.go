package network_device_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDeviceUserTagsCmd = &cobra.Command{
	Use: "list-device-user-tags [device_id]",

	Short: "Get the list of tags for a device",
	Long: `Get the list of tags for a device
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#list-device-user-tags`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListTagsResponse
		var err error

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-device-user-tags")

		fmt.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {

	Cmd.AddCommand(ListDeviceUserTagsCmd)
}
