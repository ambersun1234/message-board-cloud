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

sudo apt install protobuf-compiler -y
output=$(protoc --version)
version=$(echo ${output} | cut -d ' ' -f 2 | cut -d '.' -f 1 | bc)
if [ $version -ge 3 ]; then
    echo -e "${GREEN}Protocal buffer compiler installed successfully${NC}"
    echo -e "${output}"
else
    echo -e "${RED}Protocal buffer compiler installation failed${NC}"
    exit 1
fi