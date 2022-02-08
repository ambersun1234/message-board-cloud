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

env_file=~/.profile
folder=/tmp
filename=go1.17.6.linux-amd64.tar.gz

if ! test -f ${folder}/${filename}; then 
    cd ${folder} && wget https://go.dev/dl/${filename}
else
    echo -e "${YELLOW}Golang 1.17 download file already exists${NC}"
fi

checksum=$(sha256sum ${folder}/${filename} | cut -d ' ' -f 1)
correct_checksum="231654bbf2dab3d86c1619ce799e77b03d96f9b50770297c8f4dff8836fc8ca2"
if [ "$checksum" != "$correct_checksum" ]; then
    echo -e "${RED}Downloaded file is not correct${NC}"
    exit 1
fi

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf ${folder}/${filename}
export PATH=$PATH:/usr/local/go/bin >> ${env_file}
source ${env_file}

version=$(go version)
if [[ $version != *"1.17"* ]]; then
    echo -e "${RED}Golang 1.17 installation failed${NC}"
    exit 1
else
    echo -e "${GREEN}Golang 1.17 installed successfully${NC}"
    echo ${version}
fi