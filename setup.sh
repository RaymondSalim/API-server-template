#!/bin/sh

GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

root_dir=$(git rev-parse --show-toplevel)
cd "$root_dir" || exit

print_blue() {
  echo -e "${BLUE}$1${NC}"
}

print_yellow() {
  echo -e "${YELLOW}$1${NC}"
}

print_green() {
  echo -e "${GREEN}$1${NC}"
}

print_reminder() {
  print_green "Initialization successful. Don't forget to setup your register your remote git url"
}

reset_git() {
  rm -rf ./.git
  git init &> /dev/null
}

update_hooks_dir() {
  print_yellow "Changing git hooks path for this repository to ./.githooks"
  git config core.hooksPath .githooks
}

update_project_name() {
  print_yellow "Enter Project Name"
  read -r proj_name

  grep -rl 'API-server-template' . | xargs sed -i "s/API-server-template/${proj_name}/g"
}

reset_git
update_hooks_dir
update_project_name
print_reminder


