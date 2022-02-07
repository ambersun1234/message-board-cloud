#!/bin/sh

if ! test -d .git; then
    echo "Execute scripts/install-git-hooks in the top-level directory."
    exit 1
fi

mkdir -p .git/hooks
ln -sf ../../scripts/commit-msg.hook .git/hooks/commit-msg || exit 1
chmod +x .git/hooks/commit-msg

ln -sf ../../scripts/pre-push.hook .git/hooks/pre-push || exit 1
chmod +x .git/hooks/pre-push

touch .git/hooks/applied || exit 1

echo
echo "Git hooks are installed successfully."