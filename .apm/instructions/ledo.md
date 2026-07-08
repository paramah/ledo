---
description: Rules for editing ledo project files (.ledo.yml, compose, Taskfile, CI).
applyTo:
  - "**/.ledo.yml"
  - "**/.ledo-mode"
  - "**/docker-compose*.yml"
  - "**/Taskfile.yml"
  - "**/.gitlab-ci.yml"
  - "**/.github/workflows/*.yml"
  - "**/.gitea/workflows/*.yml"
---

# Working in a ledo (LeadDocker) project

When editing any of the files above, follow these rules. Full details are in the `ledo` skill
(`.apm/skills/ledo/SKILL.md` + `reference.md`).

- **`.ledo.yml` container block uses the `docker:` key**, not `container:`. The parser reads
  `docker:`; a `container:` block yields empty registry/namespace/name.
- **Keep modes and compose files in sync.** Every file listed in a `modes:` entry must exist,
  and the service named by `main_service` must be present in the base compose file.
- **Never add destructive ledo commands to CI or Taskfiles without a guard.** `ledo container
  down` (removes volumes), `ledo container rm`, and `ledo container prune` must not run
  unattended. For CI, use `ledo container uponce` — it runs the stack once and exits with the
  main service's exit code.
- **Only use real ledo commands/flags.** Check `.apm/skills/ledo/reference.md` before adding an
  invocation. E.g. `uponce` takes no flags; `ledo image push` takes a positional version.
- **CI pattern:** install ledo → `ledo mode select test` → `ledo container uponce` for tests;
  on tags, registry login → `ledo image build <version>` → `ledo image push <version>`.
