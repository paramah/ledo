![logo](./docs/logo.png)

# Table of contents

- [What is ledo?](#about)
- [Install](#install)
  - [Using binary](#using-binnary)
  - [Using go](#using-go-install)
- [Usage](#usage)
  - [Init](#init)
  - [Docker compose basics](#docker-compose)
- [Thanks](#thanks)

# About

Ledo (LeadDocker) is a simple tool to facilitate the daily work with docker-compose in a project (it doesn't work in swarm for now). It allows you to create run modes and fully automate them.

Ledo supports `docker` and `podman`, in the case of `podman` you need to configure it properly on your own system (links to tutorials can be found below)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/paramah/ledo)

[![Chat on Matrix](https://matrix.to/img/matrix-badge.svg)](https://matrix.to/#/#ledo:matrix.cynarski.dev)

# Install

## Using binnary

Go to [Release page](https://github.com/paramah/ledo/releases), download, unpack and move to some `PATH` directory.

You can also use the installation script:

```bash
curl -sL https://raw.githubusercontent.com/paramah/ledo/master/install.sh | sudo sh
```

## Using go install

```
go install github.com/paramah/ledo@v1.2.0
```

# Usage

## Project code structure

```
├── app (**application**)
├── docker (**docker/podman stack**)
│   ├── docker-compose.dev.yml
│   ├── docker-compose.launch.yml
│   ├── docker-compose.test.yml
│   ├── docker-compose.yml
│   ├── docker-entrypoint.sh
│   ├── etc (**files to be copied into the container**)
│   └── test-entrypoint.sh
├── Dockerfile
```

## Init

Using the `ledo init` command, you can create a `.ledo.yml` configuration file that will contain all the necessary data to use ledo, example configuration file:

```yaml
runtime: docker
docker:
  registry: registry.example.com
  namespace: Test
  name: test-application
  main_service: test-application
  shell: /bin/bash
  username: www-data
modes:
  dev: docker/docker-compose.yml docker/docker-compose.dev.yml
  traefik: docker/docker-compose.yml docker/docker-compose.traefik.yml
```

For comfortable work with `ledo` modes are important. As you can see in the example above we have defined two modes (dev and traefik).

## Docker compose

Ledo executes the `docker-compose` command with properly prepared parameters, depending on the mode selected

[![asciicast](https://asciinema.org/a/fPVl1wmtZpZXnPl3ZazoenUhD.png)](https://asciinema.org/a/fPVl1wmtZpZXnPl3ZazoenUhD)

## Podman compose

For `podman` Ledo executes `podman-compose` command.

## Podman configuration

- [Podman tutorial](https://github.com/containers/podman/blob/main/docs/tutorials/podman_tutorial.md)
- [Rootless tutorial](https://github.com/containers/podman/blob/main/docs/tutorials/rootless_tutorial.md)

# Thanks

- [Jazzy Innovations](https://jazzy.pro)
- [Stream Sage](https://streamsage.io)
