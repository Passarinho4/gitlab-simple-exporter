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
# HELP gitlab_ci_build_duration Duration in seconds of the most recent job
# TYPE gitlab_ci_build_duration gauge
gitlab_ci_build_duration{branch="master",build_stage="build",build_status="success",pipeline_id="157454",repo="https://gitlab.com/xxx"} 73.70997619628906
gitlab_ci_build_duration{branch="master",build_stage="patch-version",build_status="success",pipeline_id="157454",repo="https://gitlab.com/xxx"} 11.628866195678711
# HELP gitlab_ci_pipeline_duration Duration in seconds of the most recent pipeline
# TYPE gitlab_ci_pipeline_duration gauge
gitlab_ci_pipeline_duration{branch="master",pipeline_id="157454",repo="https://gitlab.com/xxx"} 85
# HELP gitlab_ci_pipeline_job_timestamp Creation date timestamp of the most recent pipeline
# TYPE gitlab_ci_pipeline_job_timestamp gauge
gitlab_ci_pipeline_job_timestamp{branch="master",pipeline_id="157454",repo="https://gitlab.com/xxx"} 1.700339585e+09
# HELP gitlab_ci_pipeline_total Number of executed pipelines
# TYPE gitlab_ci_pipeline_total counter
gitlab_ci_pipeline_total{branch="master",repo="https://gitlab.com/xxx",status="success"} 1
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0
```

Please note that the above example is anonymized. The metrics include data such as build durations, pipeline counters, and errors encountered by the Prometheus HTTP metric handler.

## Known Issues
Date Parsing Format: The application currently uses the "2006-01-02 15:04:05 -0700" format for parsing dates. Consider this when working with date-related functionalities.

Race condition: There might be a race condition during garbage collection job - this will be fixed in the future. But right now it is not critical in my usecase. 

## Future Enhancements
In the future, we plan to add a dashboard to provide a more user-friendly visualization of the GitLab pipeline metrics.

## Contributing
Contributions to this project are welcome! If you encounter issues or have suggestions, please create a GitHub issue. Pull requests for bug fixes, features, or improvements are also appreciated.

Thank you for using gitlab-simple-exporter!