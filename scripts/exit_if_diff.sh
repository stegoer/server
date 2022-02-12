#!/usr/bin/env sh

diff=$(git diff -- . ':(exclude)*go.sum' ':(exclude)*.sh')

if [ -n "$diff" ]; then
  echo "$diff"
  exit 1
fi
