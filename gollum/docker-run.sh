#!/bin/bash

# Initialize the wiki
if [ ! -d .git ] && [ "$(git rev-parse  --is-bare-repository 2>/dev/null)" != "true" ]; then
    git init
fi

# Set git user.name and user.email
if [ ${GOLLUM_AUTHOR_USERNAME:+1} ]; then
	git config user.name "${GOLLUM_AUTHOR_USERNAME}"
fi
if [ ${GOLLUM_AUTHOR_EMAIL:+1} ]; then
	git config user.email "${GOLLUM_AUTHOR_EMAIL}"
fi

INGRESS_URL=$(curl -H "Authorization: Bearer $SUPERVISOR_TOKEN" -s supervisor/addons/12c9acea_gollum/info | jq -r .data.ingress_url)

cp /custom.css /data/custom.css
git diff --quiet && git diff --staged --quiet || git commit -am 'update custom.css'

# Start gollum service
exec gollum --base-path $INGRESS_URL --template-dir /templates --allow-uploads page --css $@
