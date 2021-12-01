package templates

var LedoConfigurationFileTemplate = `
docker:
  registry: {{.Registry}}
  namespace: {{.Namespace}}
  name: {{.Name}}
  main_service: {{.MainService}}
  shell: {{.Shell}}
  {{- if ne .Username "root" }}
  username: {{.Username}}
  {{end}}
modes:
  base: docker/docker-compose.yml
  dev: docker/docker-compose.yml docker/docker-compose.dev.yml
  test: docker/docker-compose.yml docker/docker-compose.test.yml
`
