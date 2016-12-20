#!/bin/bash

set -e

gen() {
  (
    cd "$1" || exit
    ./generate_mocks.sh
  )
}

echo "Generating all mocks"

gen email

echo "Generating all mocks finished"
