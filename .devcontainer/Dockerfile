# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

ARG VARIANT="1.25-bookworm"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}

# [Optional] Install an additional version of node using nvm
# ARG NODE_VERSION="none"
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"

# Install Docker CLI (no daemon) and compose plugin
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends ca-certificates curl gnupg \
    && install -m 0755 -d /etc/apt/keyrings \
    && curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc \
    && chmod a+r /etc/apt/keyrings/docker.asc \
    && . /etc/os-release \
    && echo "Types: deb" > /etc/apt/sources.list.d/docker.sources \
    && echo "URIs: https://download.docker.com/linux/debian" >> /etc/apt/sources.list.d/docker.sources \
    && echo "Suites: ${VERSION_CODENAME}" >> /etc/apt/sources.list.d/docker.sources \
    && echo "Components: stable" >> /etc/apt/sources.list.d/docker.sources \
    && echo "Signed-By: /etc/apt/keyrings/docker.asc" >> /etc/apt/sources.list.d/docker.sources \
    && apt-get update \
    && apt-get -y install --no-install-recommends docker-ce-cli docker-compose-plugin docker-buildx-plugin \
    && rm -rf /var/lib/apt/lists/*

# vault-log-audit and immudb-log-audit are built from source inside the workspace.
# Run `make immudb-log-audit vault-log-audit` after the container is created.

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
