# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

ARG VARIANT="1.25-trixie"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}

# [Optional] Install an additional version of node using nvm
# ARG NODE_VERSION="none"
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"

# vault-log-audit and immudb-log-audit are built from source inside the workspace.
# Run `make immudb-log-audit vault-log-audit` after the container is created.

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
