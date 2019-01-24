## Installing using Docker

http://docs.grafana.org/installation/docker/

`$ docker run -d -p 3000:3000 grafana/grafana`

All options defined in `conf/grafana.ini` can be overridden using environment variables by using the syntax `GF_<SectionName>_<KeyName>`. For example:

```
$ docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_SERVER_ROOT_URL=http://grafana.server.name" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  grafana/grafana
```


## Using environment variables¶
http://docs.grafana.org/installation/configuration/#using-environment-variables


All options in the configuration file (listed below) can be overridden using environment variables using the syntax:

`GF_<SectionName>_<KeyName>`


Where the section name is the text within the brackets. Everything should be upper case, . should be replaced by _. For example, given these configuration settings:
```
# default section
instance_name = ${HOSTNAME}

[security]
admin_user = admin

[auth.google]
client_secret = 0ldS3cretKey
```

Then you can override them using:
```
export GF_DEFAULT_INSTANCE_NAME=my-instance
export GF_SECURITY_ADMIN_USER=true
export GF_AUTH_GOOGLE_CLIENT_SECRET=newS3cretKey
```


```
Configuration
Comments In .ini Files
Config file locations
Using environment variables
instance_name
[paths]
data
temp_data_lifetime
logs
plugins
provisioning
[server]
http_addr
http_port
protocol
socket
domain
enforce_domain
root_url
static_root_path
cert_file
cert_key
router_logging
[database]
url
type
path
host
name
user
password
ssl_mode
ca_cert_path
client_key_path
client_cert_path
server_cert_name
max_idle_conn
max_open_conn
conn_max_lifetime
log_queries
cache_mode
[security]
admin_user
admin_password
login_remember_days
secret_key
disable_gravatar
data_source_proxy_whitelist
[users]
allow_sign_up
allow_org_create
auto_assign_org
auto_assign_org_id
auto_assign_org_role
viewers_can_edit
[auth]
[session]
provider
provider_config
cookie_name
cookie_secure
session_life_time
[analytics]
reporting_enabled
google_analytics_ua_id
[dashboards]
versions_to_keep
[dashboards.json]
enabled
path
[smtp]
enabled
host
user
password
cert_file
key_file
skip_verify
from_address
from_name
ehlo_identity
[log]
mode
level
filters
[metrics]
enabled
basic_auth_username
basic_auth_password
interval_seconds
[metrics.graphite]
address
prefix
[snapshots]
external_enabled
external_snapshot_url
external_snapshot_name
snapshot_remove_expired
[external_image_storage]
provider
[external_image_storage.s3]
bucket
region
path
bucket_url
access_key
secret_key
[external_image_storage.webdav]
url
public_url
username
password
[external_image_storage.gcs]
key_file
bucket name
path
[external_image_storage.azure_blob]
account_name
account_key
container_name
[alerting]
enabled
execute_alerts
error_or_timeout
nodata_or_nullvalues
concurrent_render_limit
```