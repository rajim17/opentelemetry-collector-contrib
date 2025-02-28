# Elasticsearch Exporter

| Status                   |             |
| ------------------------ |-------------|
| Stability                | [beta]      |
| Supported pipeline types | logs,traces |
| Distributions            | [contrib]   |

This exporter supports sending OpenTelemetry logs to [Elasticsearch](https://www.elastic.co/elasticsearch).

## Configuration options

- `endpoints`: List of Elasticsearch URLs. If endpoints and cloudid is missing, the
  ELASTICSEARCH_URL environment variable will be used.
- `cloudid` (optional):
  [ID](https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html) of the
  Elastic Cloud Cluster to publish events to. The `cloudid` can be used instead
  of `endpoints`.
- `num_workers` (optional): Number of workers publishing bulk requests concurrently.
- `index`: The
  [index](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices.html)
  or [datastream](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html)
  name to publish events to. The default value is `logs-generic-default`. Note: To better differentiate between log indexes and traces indexes, `index` option are deprecated and replaced with below `logs_index`
- `logs_index`: The
  [index](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices.html)
  or [datastream](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html)
  name to publish events to. The default value is `logs-generic-default`
- `traces_index`: The
  [index](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices.html)
  or [datastream](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html)
  name to publish traces to. The default value is `traces-generic-default`.
- `pipeline` (optional): Optional [Ingest Node](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html)
  pipeline ID used for processing documents published by the exporter.
- `flush`: Event bulk buffer flush settings
  - `bytes` (default=5242880): Write buffer flush limit.
  - `interval` (default=30s): Write buffer time limit.
- `retry`: Event retry settings
  - `enabled` (default=true): Enable/Disable event retry on error. Retry
    support is enabled by default.
  - `max_requests` (default=3): Number of HTTP request retries.
  - `initial_interval` (default=100ms): Initial waiting time if a HTTP request failed.
  - `max_interval` (default=1m): Max waiting time if a HTTP request failed.
- `mapping`: Events are encoded to JSON. The `mapping` allows users to
  configure additional mapping rules.
  - `mode` (default=ecs): The fields naming mode. valid modes are:
    - `none`: Use original fields and event structure from the OTLP event.
    - `ecs`: Try to map fields defined in the
             [OpenTelemetry Semantic Conventions](https://github.com/open-telemetry/opentelemetry-specification/tree/main/semantic_conventions)
             to [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/current/index.html).
  - `fields` (optional): Configure additional fields mappings.
  - `file` (optional): Read additional field mappings from the provided YAML file.
  - `dedup` (default=true): Try to find and remove duplicate fields/attributes
    from events before publishing to Elasticsearch. Some structured logging
    libraries can produce duplicate fields (for example zap). Elasticsearch
    will reject documents that have duplicate fields.
  - `dedot` (default=true): When enabled attributes with `.` will be split into
    proper json objects.
- `sending_queue`
  - `enabled` (default = false)
  - `num_consumers` (default = 10): Number of consumers that dequeue batches; ignored if `enabled` is `false`
  - `queue_size` (default = 1000): Maximum number of batches kept in memory before data; ignored if `enabled` is `false`;
### HTTP settings

- `read_buffer_size` (default=0): Read buffer size.
- `write_buffer_size` (default=0): Write buffer size used when.
- `timeout` (default=90s): HTTP request time limit.
- `headers` (optional): Headers to be send with each HTTP request.

### Security and Authentication settings

- `user` (optional): Username used for HTTP Basic Authentication.
- `password` (optional): Password used for HTTP Basic Authentication.
- `api_key` (optional):  Authorization [API Key](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html).

### TLS settings
- `ca_file` (optional): Root Certificate Authority (CA) certificate, for
  verifying the server's identity, if TLS is enabled.
- `cert_file` (optional): Client TLS certificate.
- `key_file` (optional): Client TLS key.
- `insecure` (optional): In gRPC when set to true, this is used to disable the client transport security. In HTTP, this disables verifying the server's certificate chain and host name.
- `insecure_skip_verify` (optional): Will enable TLS but not verify the certificate.
  is enabled.

### Node Discovery

The Elasticsearch Exporter will check Elasticsearch regularly for available
nodes and updates the list of hosts if discovery is enabled. Newly discovered
nodes will automatically be used for load balancing.

- `discover`:
  - `on_start` (optional): If enabled the exporter queries Elasticsearch
    for all known nodes in the cluster on startup.
  - `interval` (optional): Interval to update the list of Elasticsearch nodes.

## Example

```yaml
exporters:
  elasticsearch/trace:
    endpoints: [https://elastic.example.com:9200]
    traces_index: trace_index
  elasticsearch/log:
    endpoints: [http://localhost:9200]
    logs_index: my_log_index
    sending_queue:
      enabled: true
      num_consumers: 20
      queue_size: 1000
······
service:
  pipelines:
    logs:
      receivers: [otlp]
      processors: [batch]
      exporters: [elasticsearch/log]
    traces:
      receivers: [otlp]
      exporters: [elasticsearch/trace]
      processors: [batch]
```
[beta]:https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]:https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib