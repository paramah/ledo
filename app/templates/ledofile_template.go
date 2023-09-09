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
  base: container/container-compose.yml
  dev: container/container-compose.yml container/container-compose.dev.yml
  test: container/container-compose.yml container/container-compose.test.yml
`
