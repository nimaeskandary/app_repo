# Using Agents

Codex should run inside a Docker Sandbox (`sbx`) for this project. The sandbox
can access the project workspace and its own Docker engine without access to
other host directories such as `~/.ssh` or `~/.aws`.

## Dependencies

- Docker Desktop
- [Docker Sandboxes (`sbx`)](https://docs.docker.com/ai/sandboxes/)
- OpenAI Codex access

Sign in with `sbx login` before creating the sandbox. Codex itself is included
in the sandbox template, so a separate host installation is not required.

## Sandbox template

`docker/codex-sbx/Dockerfile` extends Docker's Docker-enabled Codex template.
It provides the project's pinned Go, Node, and npm versions, plus a private
Docker engine for Compose services, Testcontainers, and test suites.

The current template targets Linux ARM64.

## Initial setup

Run this once from the repository root:

```sh
bin/setup-codex-sbx
```

The setup script:

1. Builds the `app_repo-codex:latest` template.
2. Loads it into the Docker Sandbox image store.
3. Creates the `app-repo` sandbox with this repository mounted read-write.
4. Prints the installed tool versions.

## Start Codex

Start or reconnect to a Codex session from the repository root:

```sh
bin/codex-sbx
```

The sandbox persists after Codex exits. Stop it when it is not needed:

```sh
sbx stop app-repo
```

Open a shell inside it for troubleshooting:

```sh
sbx exec -it app-repo bash
```

## Workspace and Docker state

The repository is directly mounted into the sandbox. File edits and deletions
made by Codex appear immediately in the host working tree and can be reviewed
with normal Git commands.

Packages installed outside the workspace, Docker images, containers, and
volumes remain inside the sandbox. Removing the sandbox deletes that internal
state but does not delete the host working tree.

To rebuild the sandbox after changing its Dockerfile:

```sh
sbx stop app-repo
sbx rm app-repo
bin/setup-codex-sbx
```
