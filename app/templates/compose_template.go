package templates

var LedoDockerComposeBaseFileTemplate_base = `version: '3.7'
services:

  {{.Config.Container.MainService}}:
    image: {{.Config.Container.Registry}}/{{.Config.Container.Namespace}}/{{.Config.Container.Name}}:latest
    networks:
      - network-{{.Config.Container.Namespace}}
    env_file: ${PWD}/.env

networks:
  network-{{.Config.Container.Namespace}}: {}
`

var LedoDockerComposeBaseFileTemplate_test = `version: '3.7'
services:

  {{.Config.Container.MainService}}:
    image: {{.Config.Container.Registry}}/{{.Config.Container.Namespace}}/{{.Config.Container.Name}}:latest
	entrypoint:
      - test-entrypoint.sh
    networks:
      - test-network-{{.Config.Container.Namespace}}
    env_file: ${PWD}/.env

networks:
  test-network-{{.Config.Container.Namespace}}: {}
`

var LedoDockerComposeBaseFileTemplate_dev = `version: '3.7'
services:

  {{.Config.Container.MainService}}:
    build:
      context: ../
      args:
        ENVIRONMENT: development
    volumes:
      - '../app:/var/www'
    ports:
      - '8090:80'
`

var LedoDockerComposeBaseFileTemplate_traefik = `version: '3.7'
services:

  {{.Config.Container.MainService}}:
    environment:
      APP_ENV: prod
    labels:
      - "traefik.enable=true"
      - "traefik.container.network=network-{{.Config.Container.Namespace}}"
      - "traefik.http.routers.{{.Config.Container.MainService}}.rule=Host(app.example.com)"
      - "traefik.http.routers.{{.Config.Container.MainService}}.service={{.Config.Container.MainService}}"
      - "traefik.http.services.{{.Config.Container.MainService}}.loadbalancer.server.port=80"
`
