#!/bin/bash

source ./tools/colors.sh

RED=
YELLOW=
BLUE=
WHITE=
CYAN=
GREEN=
NC=

set_colors

PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d ~/.local
export PATH="$PATH:~/.local/bin" >> ~/.bashrc
source ~/.bashrc

output=$(protoc --version)
version=$(echo ${output} | cut -d ' ' -f 2)
major_version=$(echo ${version} | cut -d '.' -f 1 | bc)
minor_version=$(echo ${version} | cut -d '.' -f 2 | bc)
if [ $major_version -ge 3 ] && [ $minor_version -ge 15 ]; then
    echo -e "${GREEN}Protocal buffer compiler installed successfully${NC}"
    echo -e "${output}"
else
    echo -e "${RED}Protocal buffer compiler installation failed${NC}"
    exit 1
fi