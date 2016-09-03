#!/bin/bash

case $1 in
  "unit")
    ./test-unit.sh
    ;;
  "integration")
    ./test-integration.sh
    ;;
  *)
    echo "usage: $0 [run|unit|integration]"
    exit 1
    ;;
esac
