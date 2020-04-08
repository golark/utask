#!/bin/bash
set -e
# utask installation script
#
# See https://github.com/golark/utaskd for more information
#
# This script builds the binary & installs the application
#
# Usage:
#   $ sh install_utask.sh
#
# Make sure the contentss of the script is the same as
# https://github.com/golark/utask/install.sh
# before executing

# top level definitions
TOP_DIR=$(dirname "$0")
INSTALL_DIR="/opt/utask" # final installation directory
CONFIG_FILE="./utaskcfg.yaml"
BIN_DIR="./bin" # this is where the binaries will be generated prior to installation

# distribution check
get_distribution() {
	dist=""

  # check the os-release file
	if [ -r /etc/os-release ]; then
		dist="$(. /etc/os-release && echo "$ID")"
	fi

	echo "$dist"
}

# pause for a few seconds determined by argument 1 while printing dots
sleep_with_dots() {

  for ((i=1; i<=$1; i++))
    do
      printf "."
      sleep 1
    done
  printf "\n"

}

check_prerequisites() {

  echo "checking prerequisites"

  echo "######################################################"
  # step 1 - check distribution
  dist=$(get_distribution)
  echo "detected OS: $dist"

  # step 2 - check golang installation
  ret=$(go version)
  exitCode=$?
  if [ 0 != $exitCode ]; then
    >&2 echo "can not get go version, check your installation. exiting..."
    exit $exitCode
  else
    echo "detected go: $ret"
  fi

  # step 3 - check GOPATH
  if [ -z $GOPATH ]; then
    >&2 echo "GOPATH is not set exiting"
    exit 1
  else
    echo "GOPATH: $GOPATH"
  fi

}

# try to generate binary at first argument
generate_binary() {

  OUTPUT_DIR=$1

  # check if the OUTPUT_DIR is set
  if [ -z "$OUTPUT_DIR" ]
  then
      >&2 echo "OUTPUT_DIR is empty, can not generate binary, checj arguments"
      exit 1
  fi

  echo "######################################################"

  echo "generating binary"
  # generate binary at the given directory
  rm -rf $OUTPUT_DIR && mkdir $OUTPUT_DIR
  ret=$(go build -o $OUTPUT_DIR)
}

# install application and required files
install_app() {

  # install apllication
  sudo mkdir -p $INSTALL_DIR           # create installation folder
  sudo chown $USER:$USER $INSTALL_DIR  # set permission/group to current user
  echo "installing  binary"
  cp "$BIN_DIR/utask" $INSTALL_DIR          # copy binary
  cp $CONFIG_FILE $INSTALL_DIR         # copy config file

  echo "application installed to $INSTALL_DIR"
}


# first of all check prerequisites
check_prerequisites

# generate required bins
generate_binary $BIN_DIR

# lets install
install_app

# reminder to add path
echo "** remember to add $INSTALL_DIR to your path  \$ export PATH=\$PATH:$INSTALL_DIR"
