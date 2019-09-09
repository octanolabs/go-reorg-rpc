#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
octanodir="$workspace/src/github.com/octanolabs"
if [ ! -L "$octanodir/go-reorg-rpc" ]; then
    mkdir -p "$octanodir"
    cd "$octanodir"
    ln -s ../../../../../. go-reorg-rpc
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$octanodir/go-reorg-rpc"
PWD="$octanodir/go-reorg-rpc"

# Launch the arguments with the configured environment.
exec "$@"
