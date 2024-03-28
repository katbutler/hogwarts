#!/bin/bash
# Usage: sudo setupusers.sh <names.txt>

set -xe

F=$1

if [ -z "$F" ]; then
  echo "Usage: sudo setupusers.sh <names.txt>"
  exit 1
fi

NAMES=$(cat $F)

for name in ${NAMES[@]}
do
  echo $name
  useradd -d "/home/$name" -s /bin/bash -R "/home/$name" -U -m --skel /home/hpotter "$name"  
done
