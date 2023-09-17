package templates

var LedoConfigurationFileTemplate = `
container:
  registry: {{.Registry}}
  namespace: {{.Namespace}}
  name: {{.Name}}
  main_service: {{.MainService}}
  shell: {{.Shell}}
  {{- if ne .Username "root" }}
  username: {{.Username}}
  {{end}}
modes:
  base: container/docker-compose.yml
  dev: container/docker-compose.yml container/docker-compose.dev.yml
  test: container/docker-compose.yml container/docker-compose.test.yml
`
