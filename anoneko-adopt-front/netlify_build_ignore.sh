#!/bin/bash
ls -l
if git diff HEAD^ HEAD /anoneko-adopt-front
then
  # No changes in anoneko-adopt-front, exit with 0 to skip build
  echo "No changes in anoneko-adopt-front, skipping build"
  exit 0
else
  # Changes detected, exit with 1 to proceed with build
  echo "Changes detected in anoneko-adopt-front, building"
  exit 1
fi