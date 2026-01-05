package cmd

import (
	"github.com/gplassard/woof/cmd/action_connection"
	"github.com/gplassard/woof/cmd/actions_datastores"
	"github.com/gplassard/woof/cmd/agentless_scanning"
	"github.com/gplassard/woof/cmd/api_management"
	"github.com/gplassard/woof/cmd/apm"
	"github.com/gplassard/woof/cmd/apm_retention_filters"
	"github.com/gplassard/woof/cmd/app_builder"
	"github.com/gplassard/woof/cmd/application_security"
	"github.com/gplassard/woof/cmd/audit"
	"github.com/gplassard/woof/cmd/authn_mappings"
	"github.com/gplassard/woof/cmd/aws_integration"
	"github.com/gplassard/woof/cmd/aws_logs_integration"
	"github.com/gplassard/woof/cmd/case_management"
	"github.com/gplassard/woof/cmd/case_management_attribute"
	"github.com/gplassard/woof/cmd/case_management_type"
	"github.com/gplassard/woof/cmd/ci_visibility_pipelines"
	"github.com/gplassard/woof/cmd/ci_visibility_tests"
	"github.com/gplassard/woof/cmd/cloud_cost_management"
	"github.com/gplassard/woof/cmd/cloud_network_monitoring"
	"github.com/gplassard/woof/cmd/cloudflare_integration"
	"github.com/gplassard/woof/cmd/confluent_cloud"
	"github.com/gplassard/woof/cmd/container_images"
	"github.com/gplassard/woof/cmd/containers"
	"github.com/gplassard/woof/cmd/csm_agents"
	"github.com/gplassard/woof/cmd/csm_coverage_analysis"
	"github.com/gplassard/woof/cmd/csm_threats"
	"github.com/gplassard/woof/cmd/dashboard_lists"
	"github.com/gplassard/woof/cmd/data_deletion"
	"github.com/gplassard/woof/cmd/datasets"
	"github.com/gplassard/woof/cmd/deployment_gates"
	"github.com/gplassard/woof/cmd/domain_allowlist"
	"github.com/gplassard/woof/cmd/dora_metrics"
	"github.com/gplassard/woof/cmd/downtimes"
	"github.com/gplassard/woof/cmd/error_tracking"
	"github.com/gplassard/woof/cmd/events"
	"github.com/gplassard/woof/cmd/fastly_integration"
	"github.com/gplassard/woof/cmd/fleet_automation"
	"github.com/gplassard/woof/cmd/gcp_integration"
	"github.com/gplassard/woof/cmd/incident_services"
	"github.com/gplassard/woof/cmd/incident_teams"
	"github.com/gplassard/woof/cmd/incidents"
	"github.com/gplassard/woof/cmd/ip_allowlist"
	"github.com/gplassard/woof/cmd/key_management"
	"github.com/gplassard/woof/cmd/logs"
	"github.com/gplassard/woof/cmd/logs_archives"
	"github.com/gplassard/woof/cmd/logs_custom_destinations"
	"github.com/gplassard/woof/cmd/logs_metrics"
	"github.com/gplassard/woof/cmd/logs_restriction_queries"
	"github.com/gplassard/woof/cmd/metrics"
	"github.com/gplassard/woof/cmd/microsoft_teams_integration"
	"github.com/gplassard/woof/cmd/monitors"
	"github.com/gplassard/woof/cmd/network_device_monitoring"
	"github.com/gplassard/woof/cmd/observability_pipelines"
	"github.com/gplassard/woof/cmd/okta_integration"
	"github.com/gplassard/woof/cmd/on_call"
	"github.com/gplassard/woof/cmd/on_call_paging"
	"github.com/gplassard/woof/cmd/opsgenie_integration"
	"github.com/gplassard/woof/cmd/org_connections"
	"github.com/gplassard/woof/cmd/organizations"
	"github.com/gplassard/woof/cmd/powerpack"
	"github.com/gplassard/woof/cmd/processes"
	"github.com/gplassard/woof/cmd/reference_tables"
	"github.com/gplassard/woof/cmd/restriction_policies"
	"github.com/gplassard/woof/cmd/roles"
	"github.com/gplassard/woof/cmd/rum"
	"github.com/gplassard/woof/cmd/rum_audience_management"
	"github.com/gplassard/woof/cmd/rum_metrics"
	"github.com/gplassard/woof/cmd/rum_retention_filters"
	"github.com/gplassard/woof/cmd/security_monitoring"
	"github.com/gplassard/woof/cmd/sensitive_data_scanner"
	"github.com/gplassard/woof/cmd/service_accounts"
	"github.com/gplassard/woof/cmd/service_definition"
	"github.com/gplassard/woof/cmd/service_level_objectives"
	"github.com/gplassard/woof/cmd/service_scorecards"
	"github.com/gplassard/woof/cmd/software_catalog"
	"github.com/gplassard/woof/cmd/spa"
	"github.com/gplassard/woof/cmd/spans"
	"github.com/gplassard/woof/cmd/spans_metrics"
	"github.com/gplassard/woof/cmd/static_analysis"
	"github.com/gplassard/woof/cmd/synthetics"
	"github.com/gplassard/woof/cmd/teams"
	"github.com/gplassard/woof/cmd/test_optimization"
	"github.com/gplassard/woof/cmd/usage_metering"
	"github.com/gplassard/woof/cmd/users"
	"github.com/gplassard/woof/cmd/workflow_automation"
	"github.com/gplassard/woof/pkg/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "woof",
	Short: "woof is a Datadog CLI",
	Long:  `woof is a Datadog CLI built with Cobra.`,
}

func Execute() {
	defaultSite := os.Getenv("DD_SITE")
	if defaultSite == "" {
		defaultSite = "datadoghq.com"
	}

	rootCmd.PersistentFlags().StringVar(&config.ApiKey, "api-key", os.Getenv("DD_API_KEY"), "Datadog API Key")
	rootCmd.PersistentFlags().StringVar(&config.AppKey, "app-key", os.Getenv("DD_APP_KEY"), "Datadog App Key")
	rootCmd.PersistentFlags().StringVar(&config.Site, "site", defaultSite, "Datadog Site")
	rootCmd.PersistentFlags().BoolVar(&config.Debug, "debug", false, "Enable debug logging")

	rootCmd.AddCommand(
		action_connection.Cmd,
		actions_datastores.Cmd,
		agentless_scanning.Cmd,
		api_management.Cmd,
		apm.Cmd,
		apm_retention_filters.Cmd,
		app_builder.Cmd,
		application_security.Cmd,
		audit.Cmd,
		authn_mappings.Cmd,
		aws_integration.Cmd,
		aws_logs_integration.Cmd,
		case_management.Cmd,
		case_management_attribute.Cmd,
		case_management_type.Cmd,
		ci_visibility_pipelines.Cmd,
		ci_visibility_tests.Cmd,
		cloud_cost_management.Cmd,
		cloud_network_monitoring.Cmd,
		cloudflare_integration.Cmd,
		confluent_cloud.Cmd,
		container_images.Cmd,
		containers.Cmd,
		csm_agents.Cmd,
		csm_coverage_analysis.Cmd,
		csm_threats.Cmd,
		dashboard_lists.Cmd,
		data_deletion.Cmd,
		datasets.Cmd,
		deployment_gates.Cmd,
		domain_allowlist.Cmd,
		dora_metrics.Cmd,
		downtimes.Cmd,
		error_tracking.Cmd,
		events.Cmd,
		fastly_integration.Cmd,
		fleet_automation.Cmd,
		gcp_integration.Cmd,
		incident_services.Cmd,
		incident_teams.Cmd,
		incidents.Cmd,
		ip_allowlist.Cmd,
		key_management.Cmd,
		logs.Cmd,
		logs_archives.Cmd,
		logs_custom_destinations.Cmd,
		logs_metrics.Cmd,
		logs_restriction_queries.Cmd,
		metrics.Cmd,
		microsoft_teams_integration.Cmd,
		monitors.Cmd,
		network_device_monitoring.Cmd,
		observability_pipelines.Cmd,
		okta_integration.Cmd,
		on_call.Cmd,
		on_call_paging.Cmd,
		opsgenie_integration.Cmd,
		org_connections.Cmd,
		organizations.Cmd,
		powerpack.Cmd,
		processes.Cmd,
		reference_tables.Cmd,
		restriction_policies.Cmd,
		roles.Cmd,
		rum.Cmd,
		rum_audience_management.Cmd,
		rum_metrics.Cmd,
		rum_retention_filters.Cmd,
		security_monitoring.Cmd,
		sensitive_data_scanner.Cmd,
		service_accounts.Cmd,
		service_definition.Cmd,
		service_level_objectives.Cmd,
		service_scorecards.Cmd,
		software_catalog.Cmd,
		spa.Cmd,
		spans.Cmd,
		spans_metrics.Cmd,
		static_analysis.Cmd,
		synthetics.Cmd,
		teams.Cmd,
		test_optimization.Cmd,
		usage_metering.Cmd,
		users.Cmd,
		workflow_automation.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
