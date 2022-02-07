#!/usr/bin/env bash

#
# Set colour variables if the output should be coloured.
#

set_colors() {
  local default_color=$(git config --get hooks.goodcommit.color || git config --get color.ui || echo 'auto')
  if [[ $default_color == 'true' ]] || [[ $default_color == 'always' ]] || [[ $default_color == 'auto' && -t 1 ]]; then
    RED='\033[1;31m'
    YELLOW='\033[1;33m'
    BLUE='\033[1;34m'
    WHITE='\033[1;37m'
    CYAN='\033[1;36m'
    NC='\033[0m' # No Color
  fi
}