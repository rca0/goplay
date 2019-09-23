#!/bin/bash

VERSION=1.11.13
ARCH=linux-amd64

declare -a GODEPS=(\
    "github.com/go-delve/delve/cmd/dlv" \
    "github.com/golangci/golangci-lint/cmd/golangci-lint" \
)

echo -e "Installing Golang ${VERSION}"
if [[ ! -d /usr/local/go/bin ]]; then
    echo -e "Download Golang ${VERSION}"
    curl -sL https://dl.google.com/go/go$VERSION.$ARCH.tar.gz | tar xz -C /usr/local
else
    echo -e "Golang already installed"
fi

# define vars
echo "export GOPATH=$(pwd)" >> ${HOME}/.zshrc
echo "export PATH=$PATH:$GOPATH/bin" >> ${HOME}/.zshrc

# create golang directories default
mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin

# install extension for VS Code, assume it is installed
# some deps you can install: https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on
code --install-extension ms-vscode.go

for dep in ${GODEPS[@]}; do
    printf "installing package: %s\n" "$dep"
    go get -u $dep
done
