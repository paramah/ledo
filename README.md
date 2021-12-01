# LeadDocker

```

 _                _ ___          _
| |   ___ __ _ __| |   \ ___  __| |_____ _ _
| |__/ -_) _' / _' | |) / _ \/ _| / / -_) '_|
|____\___\__,_\__,_|___/\___/\__|_\_\___|_|    dev

ledo - docker-compose and docker workflow improvement tool

 USAGE
   ledo command [subcommand] [command options] [arguments...]

 DESCRIPTION
   LeadDocker (ledo) is a simple tool for improve docker anddocker-compose workflow in your project.
   What you can do with this tool:
   => create and manage docker-compose workflow in a project
   => build docker image for project (automatic fqn and docker registry)
   => login to docker registry (with AWS ECR support)

   LeadDocker is helpful in a CI/CD.
   If you want use it as docker service try my dind image: https://hub.docker.com/r/paramah/dind

   Enjoy (-_-)

 COMMANDS
   help, h  Shows a list of commands or help for one command
   HELPERS:
     docker, d  docker helper
     image, i   docker container helper
   SETUP:
     init                           init ledo in project
     mode, m                        run mode management
     shellcompletion, autocomplete  install shell completion

 OPTIONS
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

 EXAMPLES
   ledo init                       # init ledo in your project
   ledo docker ps                  # print list of docker containers


 ABOUT
   Written & maintained by Aleksander "paramah" Cynarski

   More info about ledo on https://leaddocker.tech

   Thanks for:
    StreamSage Team        https://streamsage.io
    Jazzy Innovations Team https://jazzy.pro
```
