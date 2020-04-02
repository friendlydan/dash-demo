#!/bin/bash
export TTY

./bot_auth.sh

branch=$1

echo "Commiting Fixes"
git add -A .
git commit --author="CodeLingoBot <bot@codelingo.io>" -m "Fix Issues" # TODO automate messaged based on rule

# disable git config to get prompted
echo "Pushing Fixes Branch (Username = CodeLingoBot):"
git config --local credential.helper ""
git push clbot $branch

# return to using stored config
git config credential.helper store
