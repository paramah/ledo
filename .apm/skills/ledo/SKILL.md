---
name: ledo
description: >-
  Activate when working in a project that uses (or should use) the ledo /
  LeadDocker CLI — running containers via ledo, writing or fixing .ledo.yml,
  authoring mode-aware docker-compose files, or generating CI pipelines
  (GitHub/GitLab/Gitea) and Taskfiles that invoke ledo. Also covers ledo
  image build/push, mode switching, and Vault-backed secrets.
---

# ledo (LeadDocker)

`ledo` is a small Go CLI that wraps `docker-compose` / `podman-compose`. Instead of
memorizing long compose invocations, a project declares a `.ledo.yml` with named **modes**,
and `ledo` runs the right compose files for the active mode.

Use this skill to **use ledo in a project** and to **generate the files a ledo project needs**:
`.ledo.yml`, mode-aware `docker-compose` files, a `Taskfile.yml`, and CI pipelines that call
ledo. Copy-and-adapt starting points live in [`templates/`](templates/). The exhaustive
command + config reference is [`reference.md`](reference.md) — consult it before emitting any
`ledo …` invocation so you never invent a flag or subcommand.

## Mental model

- ledo resolves everything from the **git repo root** (it runs `git rev-parse` to find it).
  `.ledo.yml`, `.env`, `.ledo-mode`, and the compose files all sit there.
- A **mode** maps a name → a *space-separated list of compose files*, merged left-to-right.
  The active mode is stored in `.ledo-mode` (plain text, defaults to `dev` if absent).
- Every `container` subcommand ultimately runs, for the current mode:
  `docker-compose --env-file <root>/.env --project-name <namespace> -f <file1> -f <file2> … <verb>`
  (`podman-compose` when `runtime: podman`).
- The **main service** (`main_service`) is the target of `shell`, `run`, `exec`, `debug`,
  and the `--exit-code-from` service for `uponce`.

## Core workflow recipes

| Goal | Command |
| --- | --- |
| Start the stack (current mode) | `ledo container up` (`-b` build first, `-l` follow logs, `-n` no-detach) |
| One-shot run for CI (up, wait, exit with main service's code) | `ledo container uponce` |
| Open a shell in the main service | `ledo container shell` (alias `sh`; `--user <name>`) |
| Run a command in a new container | `ledo container run -- <cmd…>` |
| Exec into the running main container | `ledo container exec -- <cmd…>` |
| Follow logs | `ledo container logs [service…]` (last 100 lines, `--follow`) |
| Build compose images | `ledo container build` (`-n` no-cache) |
| Stop / restart / remove | `ledo container stop` / `restart` / `rm` |
| Build the project image | `ledo image build [version] [-f Dockerfile] [-s stage]` (default tag `latest`) |
| Print image FQN | `ledo image fqn` → `registry/namespace/name` |
| Push image | `ledo image push [version]` (positional version, default `latest`) |
| Re-tag image | `ledo image retag <from> <to>` |
| List / switch modes | `ledo mode list` · `ledo mode select <mode>` |
| Read/write secrets (Vault) | `ledo secrets read` · `ledo secrets write KEY=VALUE …` |
| Login to AWS ECR | `ledo container login ecr -r <region> -k <key> -s <secret>` |
| Scaffold a new project interactively | `ledo init` |

The FQN is always `registry/namespace/name` (lowercased) from `.ledo.yml`.

## Generation tasks (the primary use case)

When asked to set up or extend a ledo project, produce these artifacts by adapting the files in
[`templates/`](templates/). Keep them **consistent with each other**: the mode names in
`.ledo.yml` must match the compose files that exist, and Taskfile/CI must only call real ledo
commands.

1. **`.ledo.yml`** — start from [`templates/ledo.yml`](templates/ledo.yml). Fill in
   `registry`, `namespace`, `name`, `main_service`. Each `modes:` entry lists the compose
   files for that mode. ⚠️ Put the container block under the **`docker:`** key (see guardrails).
2. **Mode-aware compose files** — from [`templates/docker-compose.yml`](templates/docker-compose.yml)
   (base), [`.dev.yml`](templates/docker-compose.dev.yml) (adds `build:` + source volume mount),
   and [`.test.yml`](templates/docker-compose.test.yml) (test entrypoint). One `-f` file per
   mode entry. The base service image should be `${REGISTRY}/${NAMESPACE}/${NAME}` and use
   `env_file: .env`.
3. **`Taskfile.yml`** — from [`templates/Taskfile.yml`](templates/Taskfile.yml). Each task
   delegates to ledo (`task up` → `ledo container up`, `task ci` → `ledo container uponce`,
   etc.) so contributors don't need to know ledo's flags.
4. **CI pipelines** — pick the platform and adapt:
   [`github-actions.ci.yml`](templates/github-actions.ci.yml),
   [`gitlab-ci.yml`](templates/gitlab-ci.yml), or
   [`gitea-actions.ci.yml`](templates/gitea-actions.ci.yml). The pattern is the same on all
   three: install ledo → registry login → `ledo image build` → `ledo image push` → optionally
   `ledo container uponce` to run the test suite and surface its exit code.

**Registry login in pipelines.** ledo does **not** log in to a registry for you before
`image push`/`container pull` — in this version `ledo container login` supports **AWS ECR only**
(`ledo container login ecr -r <region> -k <key> -s <secret>`). For every other registry
(Docker Hub, GitHub GHCR, GitLab, Gitea, Harbor, a private registry, …) you must run a plain
Docker/Podman login yourself first, e.g.:

```bash
echo "$REGISTRY_PASSWORD" | docker login "$(ledo image fqn | cut -d/ -f1)" -u "$REGISTRY_USERNAME" --password-stdin
```

`ledo image fqn | cut -d/ -f1` yields the registry host from `.ledo.yml`. Use `podman login`
when `runtime: podman`. Only reach for `ledo container login ecr` when the registry is ECR.

## Guardrails (verified against ledo's source)

- **Config key is `docker:`.** The parser reads the container block under `docker:`
  (`app/modules/config/ledofile.go`), but `ledo init`'s current template emits `container:` —
  a known mismatch. When authoring `.ledo.yml` by hand, use `docker:`. If a user's registry/
  namespace/name comes back empty after `ledo init`, rename their top block from `container:`
  to `docker:`.
- **Destructive commands — never run without explicit user confirmation, and never in CI
  except `uponce`:**
  - `ledo container down` runs `down --volumes` — **deletes named volumes / data**.
  - `ledo container rm` runs `rm -f`.
  - `ledo container prune` wipes **all** containers, images, volumes, and networks
    (`system prune --all --volumes --force`); it only runs in `dev` mode and prompts first.
- **User wrapping:** if `docker.username` is set and not `root`, `run`/`exec` wrap the command
  in `sudo -E -u <username>`. `main_service` and `shell` can be overridden at runtime by env
  vars `MAIN_SERVICE` / `MAIN_SHELL`.
- **Prereqs:** `docker-compose >= 1.28.0` (or `podman-compose` when `runtime: podman`).
- **Secrets** live in Vault KV2 at `/environment/data/<namespace>/<name>/<mode>` and are
  therefore **mode-scoped** — `ledo secrets read` on the `dev` mode returns different values
  than on `test`. Needs `VAULT_ADDR` + `VAULT_TOKEN` (flags `-a`/`-t` or env).

See [`reference.md`](reference.md) for the full command/flag matrix and the complete
`.ledo.yml` schema.
