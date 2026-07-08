# AGENTS.md — ledo (LeadDocker)

Guidance for AI coding agents working in or with this repo. This project is `ledo`, a Go CLI
that wraps `docker-compose` / `podman-compose` around a `.ledo.yml` config with named **modes**.

An installable agent **skill** ships in this repo (Microsoft APM package + native Claude skill):

- **Skill:** [`.apm/skills/ledo/SKILL.md`](.apm/skills/ledo/SKILL.md) — how to use ledo and how
  to generate a project's `.ledo.yml`, mode-aware compose files, `Taskfile.yml`, and CI
  pipelines (GitHub/GitLab/Gitea). Copy-and-adapt starters in
  [`.apm/skills/ledo/templates/`](.apm/skills/ledo/templates/).
- **Reference:** [`.apm/skills/ledo/reference.md`](.apm/skills/ledo/reference.md) — the exact
  command/flag matrix and `.ledo.yml` schema. Consult it before emitting any `ledo …` command.
- **Manifest:** [`apm.yml`](apm.yml). Deploy with `apm install --target claude|copilot`
  (writes `.claude/…` / `.github/copilot-instructions.md`); regenerate rather than editing
  deployed files by hand.

## Golden-path commands

```
ledo container up [-b|-l|-n]   # start the stack for the current mode
ledo container uponce          # run once, exit with main service's code (use in CI)
ledo container shell           # shell into the main service
ledo image build [version]     # build project image (FQN = registry/namespace/name)
ledo image push  [version]     # push to registry
ledo mode list | select <m>    # list / switch modes (stored in .ledo-mode, default dev)
ledo secrets read | write      # Vault KV2, scoped per mode
```

## Must-know guardrails

- `.ledo.yml`'s container block uses the **`docker:`** key (not `container:`).
- **Destructive** — require explicit confirmation, never run in CI: `ledo container down`
  (removes volumes), `ledo container rm`, `ledo container prune` (dev-mode-only, wipes
  everything).
- Requires `docker-compose >= 1.28.0` (or `podman-compose`). Paths resolve from the git root.

## Developing ledo itself (this Go repo)

- Build: `go build ./...`  ·  Test: `go test ./...`  ·  Vendored deps (`vendor/`), Go 1.24.x.
- Commands live under `app/cmd/**`; behavior under `app/modules/**`; init templates under
  `app/templates/**`. Releases via GoReleaser on `v*` tags (CircleCI).
