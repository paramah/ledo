# ledo — command & config reference

Derived from ledo's source (`main.go`, `app/cmd/**`, `app/modules/**`). Use this as the source
of truth for exact commands, aliases, and flags — do not invent options that are not listed.

## Global

- Built on `urfave/cli/v2`. No app-level flags beyond `--help/-h` and `--version/-v`.
- Every command bootstraps via `context.InitCommand`: finds the git repo root, loads
  `.ledo.yml` (fatal "Please run ledo init" if missing), reads the active mode from
  `.ledo-mode` (default `dev`), and builds the compose args
  `--env-file <root>/.env --project-name <namespace-lowercased> -f <each mode compose file>`.
- **jzcli compat:** if `.jz-project.yml` exists it is used instead of `.ledo.yml`, and
  `.jz-mode` instead of `.ledo-mode`.
- Top-level commands: `init`, `container` (aliases `c`/`docker`/`d`), `image` (`i`),
  `secrets` (`s`), `mode` (`m`), `shellcompletion`/`autocomplete`.

## `ledo init`
Interactive scaffold. Prompts for registry/namespace/name/main_service/shell/username, writes
`./.ledo.yml`; optionally builds advanced container config: `./Dockerfile`, `./docker/` with
`docker-compose.yml` + `.dev.yml` + `.test.yml`, and `docker-entrypoint.sh` / `test-entrypoint.sh`.
Makes network calls to Docker Hub to pick a base image tag. No flags.

## `ledo container` (aliases `c`, `docker`, `d`)
Runs compose for the current mode. `before` hook checks `docker-compose >= 1.28.0` (or that
`podman-compose` is installed). Underlying invocation shown per subcommand.

| Subcommand | Aliases | Flags / args | Effect (compose verb) |
| --- | --- | --- | --- |
| `up` | `u` | `--no-detach/-n`, `--build/-b`, `--logs/-l` | `up [-d]`; optional build first, logs after |
| `uponce` | — | none | `up --force-recreate --renew-anon-volumes --abort-on-container-exit --exit-code-from <main>` |
| `build` | `b` | `--no-cache/-n` | `build --pull [--no-cache]` |
| `debug` | — | none | `run --entrypoint= <main> <shell>` (entrypoint bypassed) |
| `down` | — | none | `down --volumes` ⚠️ **removes volumes** |
| `logs` | — | `[services…]` | `logs --follow --tail 100 [services]` |
| `restart` | — | none | `restart` |
| `run` | — | `<cmd…>` (≥1 arg) | `run <main> [sudo -E -u <user>] <cmd>` |
| `exec` | `e` | `<cmd…>` (≥1 arg) | `exec <main> [sudo -E -u <user>] <cmd>` |
| `rm` | — | none | `rm -f` ⚠️ |
| `shell` | `sh` | `--user/-u <name|uid>` | `exec [--user X] <main> <shell>` |
| `start` | — | none | `start` |
| `stop` | — | none | `stop` |
| `pull` | — | none | `pull` (network) |
| `ps` | `p` | none | `ps` |
| `login ecr` | `e` | `--region/-r` (env `AWS_REGION`), `--key/-k` (env `AWS_ACCESS_KEY_ID`), `--secret/-s` (env `AWS_SECRET_ACCESS_KEY`) — all required | AWS ECR `GetAuthorizationToken` → `docker/podman login` |
| `prune` | — | none | Prompts, then removes ALL containers/images/volumes/networks + `system prune --all --volumes --force`. **`dev` mode only.** ⚠️ |

Note: `run`/`exec` take the command to run; pass it after `--` to avoid flag parsing issues,
e.g. `ledo container exec -- php bin/console cache:clear`.

Note on `login`: the **only** registry login ledo implements in this version is **`ecr`**
(`ledo container login ecr`). There is no generic `ledo container login <registry>`. For any
non-ECR registry, log in with the native client before `image push` / `container pull`, e.g.
`echo "$PASS" | docker login <registry> -u <user> --password-stdin` (`podman login` under podman).

## `ledo image` (alias `i`)
FQN is built as `registry/namespace/name` (lowercased).

| Subcommand | Aliases | Flags / args | Effect |
| --- | --- | --- | --- |
| `fqn` | `f` | none | prints `registry/namespace/name` |
| `build` | `b` | positional `version` (default `latest`); `--stage/-s`, `--dockerfile/-f` (default `./Dockerfile`), `--opts/-o` (default `--compress`) | `docker/podman build -t <fqn>:<ver> -f <dockerfile> <opts> [--target <stage>] .` |
| `push` | `p` | positional `version` (default `latest`) | `docker/podman push <fqn>:<version>` |
| `retag` | `r` | `<fromTag> <toTag>` | `docker/podman tag <fqn>:<from> <fqn>:<to>` |

## `ledo secrets` (alias `s`)
HashiCorp Vault KV2 under `/environment/data/<namespace>/<name>/<mode>` (mode-scoped).

| Subcommand | Aliases | Flags / args |
| --- | --- | --- |
| `read` | `r` | `--addr/-a` (env `VAULT_ADDR`), `--token/-t` (env `VAULT_TOKEN`), `--debug/-d` — prints `KEY=VALUE` |
| `write` | `w` | `--addr/-a`, `--token/-t`, `--debug/-d`, `--input/-i <env-file>`, or positional `KEY=VALUE…` |

## `ledo mode` (alias `m`)
Bare `ledo mode` opens an interactive picker. Active mode stored in `.ledo-mode` (default `dev`).

| Subcommand | Flags / args | Effect |
| --- | --- | --- |
| `select` | `[<mode>]` (direct if given, else picker) | writes `.ledo-mode` |
| `list` | none | lists modes from `.ledo.yml` |

## `ledo shellcompletion` (alias `autocomplete`)
Arg `<bash|zsh|powershell|fish>`, flag `--install`. Downloads/writes a completion script into
the XDG config dir; with `--install` also appends to the shell rc.

## `.ledo.yml` schema

```yaml
runtime: docker            # docker | podman (default docker) -> docker-compose vs podman-compose
docker:                    # NOTE: YAML key is `docker:` (parser reads this key)
  registry: registry.example.com
  namespace: my-group      # also used as compose --project-name and network name
  name: my-app
  main_service: app        # target of shell/run/exec/debug; env override MAIN_SERVICE
  shell: /bin/bash         # default shell; env override MAIN_SHELL
  username: www-data       # if set and != root -> run/exec wrapped in `sudo -E -u <user>`
modes:                     # name -> space-separated compose files, merged in order
  base: docker/docker-compose.yml
  dev:  docker/docker-compose.yml docker/docker-compose.dev.yml
  test: docker/docker-compose.yml docker/docker-compose.test.yml
project: my-app            # optional, largely unused
deployment:                # optional
  - host: example.com
    is_secure: true
    tls_directory: /etc/certs
    mode: dev
```

### Companion files at repo root
- `.ledo-mode` — current mode (plain text). Written by `ledo mode select`.
- `.env` — passed to compose as `--env-file`.
- `.jz-project.yml` / `.jz-mode` — legacy jzcli names, used if present instead of the `.ledo.*`.

### Env-var overrides
`MAIN_SERVICE`, `MAIN_SHELL`, `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`,
`VAULT_ADDR`, `VAULT_TOKEN`.

### Known discrepancy
`ledo init`'s template writes the container block under `container:`, but the parser reads
`docker:`. Authoring by hand → use `docker:`. After `ledo init`, rename `container:` → `docker:`
if config values come back empty.
