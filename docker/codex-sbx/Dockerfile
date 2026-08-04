FROM docker.io/docker/sandbox-templates:codex-docker

ARG NODE_VERSION=26.5.0
ARG NPM_VERSION=11.17.0
ARG GO_VERSION=1.27rc2

ARG PROJECT_ROOT_DIR

ENV IN_SANDBOX=1
ENV PATH="/home/agent/go/bin:${PATH}"

# we don't want the hosts node_modules, its for a different arch
VOLUME ${PROJECT_ROOT_DIR}/node_modules
VOLUME /home/agent/.cache/go-build
VOLUME /home/agent/go/pkg/mod

USER root

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        build-essential \
        libatomic1 \
        libgtk-4-dev \
        libwebkitgtk-6.0-dev \
        pkg-config \
        xz-utils \
    && rm -rf /var/lib/apt/lists/* \
    && codex update

RUN set -eux; \
    node_archive="node-v${NODE_VERSION}-linux-arm64.tar.xz"; \
    node_setup_dir="$(mktemp -d)"; \
    cd "$node_setup_dir"; \
    curl -fsSLO "https://nodejs.org/dist/v${NODE_VERSION}/${node_archive}"; \
    curl -fsSLO "https://nodejs.org/dist/v${NODE_VERSION}/SHASUMS256.txt"; \
    grep " ${node_archive}$" SHASUMS256.txt | sha256sum --check; \
    tar --extract --xz --file "$node_archive" --directory /usr/local --strip-components=1; \
    cd /; \
    rm -rf "$node_setup_dir"; \
    node --version; \
    npm --version; \
    test "$(npm --version)" = "$NPM_VERSION"

USER agent

RUN GOTOOLCHAIN="go${GO_VERSION}" go version
