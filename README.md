
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

# Install

## Using binnary

Go to [Release page](https://github.com/paramah/ledo/releases), download, unpack and move to some `PATH` directory. 

You can also use the installation script:

```bash
curl -sL https://raw.githubusercontent.com/paramah/ledo/master/install.sh | sudo sh
```

## Using go install

```
go install github.com/paramah/ledo@v1.0.0
```

# Usage

## Init

Using the `ledo init` command, you can create a `.ledo.yml` configuration file that will contain all the necessary data to use ledo, example configuration file:

```yaml
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

For comfortable operation modes are important, as you can see in the example above we have defined two modes (dev and traefik).

## Docker compose

Ledo executes the `docker-compose` command with properly prepared parameters, depending on the mode selected

[![asciicast](https://asciinema.org/a/fPVl1wmtZpZXnPl3ZazoenUhD.png)](https://asciinema.org/a/fPVl1wmtZpZXnPl3ZazoenUhD)

# Thanks

- [Jazzy Innovations](https://jazzy.pro)
- [Stream Sage](https://streamsage.io)
