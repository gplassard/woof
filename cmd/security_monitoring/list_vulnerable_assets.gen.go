package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use: "list-vulnerable-assets",

	Short: "List vulnerable assets",
	Long: `List vulnerable assets
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-vulnerable-assets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListVulnerableAssetsResponse
		var err error

		optionalParams := datadogV2.NewListVulnerableAssetsOptionalParameters()

		if cmd.Flags().Changed("page-token") {
			val, _ := cmd.Flags().GetString("page-token")
			optionalParams.WithPageToken(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("filter-name") {
			val, _ := cmd.Flags().GetString("filter-name")
			optionalParams.WithFilterName(val)
		}

		if cmd.Flags().Changed("filter-type") {
			val, _ := cmd.Flags().GetString("filter-type")
			optionalParams.WithFilterType(val)
		}

		if cmd.Flags().Changed("filter-version.first") {
			val, _ := cmd.Flags().GetString("filter-version.first")
			optionalParams.WithFilterVersionFirst(val)
		}

		if cmd.Flags().Changed("filter-version.last") {
			val, _ := cmd.Flags().GetString("filter-version.last")
			optionalParams.WithFilterVersionLast(val)
		}

		if cmd.Flags().Changed("filter-repository-url") {
			val, _ := cmd.Flags().GetString("filter-repository-url")
			optionalParams.WithFilterRepositoryUrl(val)
		}

		if cmd.Flags().Changed("filter-risks.in-production") {
			val, _ := cmd.Flags().GetString("filter-risks.in-production")
			optionalParams.WithFilterRisksInProduction(val)
		}

		if cmd.Flags().Changed("filter-risks.under-attack") {
			val, _ := cmd.Flags().GetString("filter-risks.under-attack")
			optionalParams.WithFilterRisksUnderAttack(val)
		}

		if cmd.Flags().Changed("filter-risks.is-publicly-accessible") {
			val, _ := cmd.Flags().GetString("filter-risks.is-publicly-accessible")
			optionalParams.WithFilterRisksIsPubliclyAccessible(val)
		}

		if cmd.Flags().Changed("filter-risks.has-privileged-access") {
			val, _ := cmd.Flags().GetString("filter-risks.has-privileged-access")
			optionalParams.WithFilterRisksHasPrivilegedAccess(val)
		}

		if cmd.Flags().Changed("filter-risks.has-access-to-sensitive-data") {
			val, _ := cmd.Flags().GetString("filter-risks.has-access-to-sensitive-data")
			optionalParams.WithFilterRisksHasAccessToSensitiveData(val)
		}

		if cmd.Flags().Changed("filter-environments") {
			val, _ := cmd.Flags().GetString("filter-environments")
			optionalParams.WithFilterEnvironments(val)
		}

		if cmd.Flags().Changed("filter-teams") {
			val, _ := cmd.Flags().GetString("filter-teams")
			optionalParams.WithFilterTeams(val)
		}

		if cmd.Flags().Changed("filter-arch") {
			val, _ := cmd.Flags().GetString("filter-arch")
			optionalParams.WithFilterArch(val)
		}

		if cmd.Flags().Changed("filter-operating-system.name") {
			val, _ := cmd.Flags().GetString("filter-operating-system.name")
			optionalParams.WithFilterOperatingSystemName(val)
		}

		if cmd.Flags().Changed("filter-operating-system.version") {
			val, _ := cmd.Flags().GetString("filter-operating-system.version")
			optionalParams.WithFilterOperatingSystemVersion(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListVulnerableAssets(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-vulnerable-assets")

		cmd.Println(cmdutil.FormatJSON(res, "assets"))
	},
}

func init() {

	ListVulnerableAssetsCmd.Flags().String("page-token", "", "Its value must come from the 'links' section of the response of the first request. Do not manually edit it.")

	ListVulnerableAssetsCmd.Flags().Int64("page-number", 0, "The page number to be retrieved. It should be equal or greater than '1'")

	ListVulnerableAssetsCmd.Flags().String("filter-name", "", "Filter by name. This field supports the usage of wildcards (*).")

	ListVulnerableAssetsCmd.Flags().String("filter-type", "", "Filter by type.")

	ListVulnerableAssetsCmd.Flags().String("filter-version.first", "", "Filter by the first version of the asset since it has been vulnerable.")

	ListVulnerableAssetsCmd.Flags().String("filter-version.last", "", "Filter by the last detected version of the asset.")

	ListVulnerableAssetsCmd.Flags().String("filter-repository-url", "", "Filter by the repository url associated to the asset.")

	ListVulnerableAssetsCmd.Flags().String("filter-risks.in-production", "", "Filter whether the asset is in production or not.")

	ListVulnerableAssetsCmd.Flags().String("filter-risks.under-attack", "", "Filter whether the asset (Service) is under attack or not.")

	ListVulnerableAssetsCmd.Flags().String("filter-risks.is-publicly-accessible", "", "Filter whether the asset (Host) is publicly accessible or not.")

	ListVulnerableAssetsCmd.Flags().String("filter-risks.has-privileged-access", "", "Filter whether the asset (Host) has privileged access or not.")

	ListVulnerableAssetsCmd.Flags().String("filter-risks.has-access-to-sensitive-data", "", "Filter whether the asset (Host)  has access to sensitive data or not.")

	ListVulnerableAssetsCmd.Flags().String("filter-environments", "", "Filter by environment.")

	ListVulnerableAssetsCmd.Flags().String("filter-teams", "", "Filter by teams.")

	ListVulnerableAssetsCmd.Flags().String("filter-arch", "", "Filter by architecture.")

	ListVulnerableAssetsCmd.Flags().String("filter-operating-system.name", "", "Filter by operating system name.")

	ListVulnerableAssetsCmd.Flags().String("filter-operating-system.version", "", "Filter by operating system version.")

	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
