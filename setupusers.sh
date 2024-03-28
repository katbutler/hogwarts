#!/bin/bash
# Usage: sudo setupusers.sh <names.txt>

set -e

F=$1

if [ -z "$F" ]; then
  echo "Usage: sudo setupusers.sh <names.txt>"
  exit 1
fi

NAMES=$(cat $F)

for name in ${NAMES[@]}
do
  echo "adding $name ..."
  set +e
  if id "$name" >/dev/null 2>&1; then
    echo "$name already exists. Skipping"
    continue
  fi
  set -e
  useradd -d "/home/$name" -s /bin/bash -U -m --skel /home/hpotter "$name"  
  echo "added $name"
done

echo "Done creating all users"
