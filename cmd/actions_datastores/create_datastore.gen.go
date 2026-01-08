package actions_datastores

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDatastoreCmd = &cobra.Command{
	Use:     "create-datastore",
	Aliases: []string{"create"},
	Short:   "Create datastore",
	Long: `Create datastore
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#create-datastore`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateAppsDatastoreResponse
		var err error

		var body datadogV2.CreateAppsDatastoreRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDatastore(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-datastore")

		fmt.Println(cmdutil.FormatJSON(res, "datastore"))
	},
}

func init() {

	CreateDatastoreCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDatastoreCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDatastoreCmd)
}
