#!/bin/bash
set -e

if date -d $1 +"%s"; then
    exit 0
else
    exit 1
fi
