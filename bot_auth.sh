#!/bin/bash
export TTY
set -x

bot_name="CodeLingo Bot"
bot_email="bot@codelingo.io"

export GIT_AUTHOR_NAME="$bot_name"
export GIT_AUTHOR_EMAIL="$bot_email"
export GIT_COMMITTER_NAME="$bot_name"
export GIT_COMMITTER_EMAIL="$bot_email"

# git config user.name "CodeLingoBot"
# git config user.email "bot@codelingo.io"