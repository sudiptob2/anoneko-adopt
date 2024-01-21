#!/bin/bash
if git diff --quiet HEAD^ HEAD anoneko-adopt-front
then
  # No changes in anoneko-adopt-front, exit with 0 to skip build
  exit 0
else
  # Changes detected, exit with 1 to proceed with build
  exit 1
fi