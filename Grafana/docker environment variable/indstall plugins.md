Installing Plugins for Grafana

http://docs.grafana.org/installation/docker/#installing-plugins-for-grafana

Pass the plugins you want installed to docker with the GF_INSTALL_PLUGINS environment variable as a comma separated list. This will pass each plugin name to grafana-cli plugins install ${plugin} and install them when Grafana starts.

```
docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_INSTALL_PLUGINS=grafana-clock-panel,grafana-simple-json-datasource" \
  grafana/grafana
```

## common plugins
grafana-clock-panel

grafana-simple-json-datasource

grafana-kubernetes-app