#!/bin/bash

set -eou pipefail

image=permission-manager
tag=$SOURCE_TAG

docker build --tag "ghcr.io/kazanexpress/$image:$tag" --tag "ghcr.io/kazanexpress/$image:latest" .
docker push --all-tags ghcr.io/kazanexpress/$image

