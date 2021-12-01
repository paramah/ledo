package templates

var LedoDockerComposeBaseFileTemplate_base = `
version: '3.7'
services:

  {{.Config.Docker.MainService}}:
    image: {{.Config.Docker.Registry}}
    networks:
      - network-{{.Config.Docker.Namespace}}
    env_file: ${PWD}/.env

networks:
  nestork-{{.Config.Docker.Namespace}}: {}
`
