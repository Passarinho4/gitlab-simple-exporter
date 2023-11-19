# gitlab-simple-exporter

**gitlab-simple-exporter** is an open-source project designed to expose GitLab pipeline data via HTTP endpoints for Prometheus consumption. It provides two HTTP endpoints:

1. `/webhook`: This endpoint is designed to receive GitLab pipeline events via webhooks. It extracts simple data from these events, including pipeline finish status, duration, and details about individual builds in the pipeline. For each build, the exporter collects information such as build duration and finish status.

2. `/metrics`: This endpoint exposes the collected data in a format compatible with Prometheus. Prometheus can scrape this endpoint to gather metrics over time.

## Usage

### Webhook Endpoint (/webhook)

To use the `/webhook` endpoint, configure your GitLab project to send webhook events for pipeline activities to the URL provided by your deployed gitlab-simple-exporter instance (e.g., `http://your-exporter-host/webhook`). The exporter will process the events and collect relevant data.

### Metrics Endpoint (/metrics)

The `/metrics` endpoint provides a Prometheus-compatible interface for retrieving the collected data. Configure Prometheus to scrape this endpoint to gather and store metrics over time.

## Example Metrics

Here is an example of the metrics exposed by the `/metrics` endpoint:

```plaintext
# HELP builds_durations Duration of build in pipeline execution
# TYPE builds_durations gauge
builds_durations{branch="master",build_stage="build",namespace="example-group",pipeline_id="123456",project_name="example-project"} 221.12
builds_durations{branch="master",build_stage="patch-version",namespace="example-group",pipeline_id="123456",project_name="example-project"} 34.89
# HELP pipeline_counter Number of executed pipelines
# TYPE pipeline_counter counter
pipeline_counter{branch="master",namespace="example-group",project_name="example-project"} 3
# HELP pipelines_durations Duration of pipeline execution
# TYPE pipelines_durations gauge
pipelines_durations{branch="master",namespace="example-group",pipeline_id="123456",project_name="example-project"} 255
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0
# HELP success_pipeline_counter Number of successfully executed pipelines
# TYPE success_pipeline_counter counter
success_pipeline_counter{branch="master",namespace="example-group",project_name="example-project"} 3
```

Please note that the above example is anonymized. The metrics include data such as build durations, pipeline counters, and errors encountered by the Prometheus HTTP metric handler.

## Known Issues
Date Parsing Format: The application currently uses the "2006-01-02 15:04:05 -0700" format for parsing dates. Consider this when working with date-related functionalities.

Metrics Reset: The application resets metrics once per hour. Ensure that Prometheus gathers newly added values before garbage collection to avoid potential data loss.

## Future Enhancements
In the future, we plan to add a dashboard to provide a more user-friendly visualization of the GitLab pipeline metrics.

## Contributing
Contributions to this project are welcome! If you encounter issues or have suggestions, please create a GitHub issue. Pull requests for bug fixes, features, or improvements are also appreciated.

Thank you for using gitlab-simple-exporter!