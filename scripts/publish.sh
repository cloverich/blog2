#!/usr/bin/env bash
# Put posts that Chronicles wrote into content/blog live: shrink their images,
# commit, and push to master (which deploys). Run from anywhere.
#
#   scripts/publish.sh            resize, commit, push
#   scripts/publish.sh --dry-run  show what would be committed, change nothing
set -euo pipefail

cd "$(dirname "$0")/.."
dry_run=false
[[ "${1:-}" == "--dry-run" ]] && dry_run=true

if [[ "$(git branch --show-current)" != "master" ]]; then
  echo "Not on master; deploys come from master." >&2
  exit 1
fi

# Posts from the iPhone arrive through the outbox (see scripts/outbox.sh).
outbox="${BLOG_OUTBOX:-$HOME/Library/Mobile Documents/com~apple~CloudDocs/Chronicles Blog}"
if [[ -d "$outbox" ]]; then
  if $dry_run; then
    echo "(dry run: not copying from the outbox; run scripts/outbox.sh to preview those posts)"
  else
    scripts/outbox.sh
  fi
fi

changes=$(git status --porcelain --untracked-files=all -- content/blog)
if [[ -z "$changes" ]]; then
  echo "Nothing to publish in content/blog."
  exit 0
fi

# Post folders (content/blog/<date>/<slug>) with any change.
posts=$(echo "$changes" | sed -E 's/^...//; s#^(content/blog/[^/]+/[^/]+).*#\1#' | sort -u)
echo "Posts:"
echo "$posts" | sed 's/^/  /'

if $dry_run; then
  go -C scripts run .
  exit 0
fi

git pull --ff-only
go -C scripts run . -write
git add -- content/blog

# "title: ..." is a JSON string in Chronicles posts and plain YAML in older
# ones; strip the quotes either way.
titles=$(for p in $posts; do
  [[ -f "$p/index.md" ]] && sed -n 's/^title: *"\{0,1\}\(.*[^"]\)"\{0,1\}$/\1/p' "$p/index.md" | head -1
done)
count=$(echo "$titles" | grep -c . || true)
if [[ "$count" == 1 ]]; then
  subject="Publish \"$titles\""
else
  subject="Publish $count posts"
fi

git commit -m "$subject" -m "$(echo "$titles" | sed 's/^/- /')"
git push
echo "Pushed; the site deploys in a minute or two."
