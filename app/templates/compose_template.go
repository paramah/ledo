package templates

var LedoDockerComposeBaseFileTemplate_base = `version: '3.7'
services:

  {{.Config.Docker.MainService}}:
    image: {{.Config.Docker.Registry}}/{{.Config.Docker.Namespace}}/{{.Config.Docker.Name}}:latest
    networks:
      - network-{{.Config.Docker.Namespace}}
    env_file: ${PWD}/.env

networks:
  network-{{.Config.Docker.Namespace}}: {}
`

var LedoDockerComposeBaseFileTemplate_test = `version: '3.7'
services:

  {{.Config.Docker.MainService}}:
    image: {{.Config.Docker.Registry}}/{{.Config.Docker.Namespace}}/{{.Config.Docker.Name}}:latest
	entrypoint:
      - test-entrypoint.sh
    networks:
      - test-network-{{.Config.Docker.Namespace}}
    env_file: ${PWD}/.env

networks:
  test-network-{{.Config.Docker.Namespace}}: {}
`

var LedoDockerComposeBaseFileTemplate_dev = `version: '3.7'
services:

  {{.Config.Docker.MainService}}:
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

  {{.Config.Docker.MainService}}:
    environment:
      APP_ENV: prod
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=network-{{.Config.Docker.Namespace}}"
      - "traefik.http.routers.{{.Config.Docker.MainService}}.rule=Host(app.example.com)"
      - "traefik.http.routers.{{.Config.Docker.MainService}}.service={{.Config.Docker.MainService}}"
      - "traefik.http.services.{{.Config.Docker.MainService}}.loadbalancer.server.port=80"
`