#!/bin/bash

set -Eex

export PATH="${PWD}:${PWD}/cached-deps:${GOPATH}/bin:${PATH}"

# Try to connect for three minutes
for _ in $(seq 36); do
    if kubectl version &>/dev/null; then
        echo 'Minikube ready!' | ts
    exit 0
    fi
    echo 'sleeping' | ts
  sleep 5
done

# Give up--kubernetes isn't coming up
exit 1
