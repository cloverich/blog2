#!/usr/bin/env bash
# Copy posts that Chronicles on the iPhone wrote to its blog folder (an iCloud
# Drive "outbox") into content/blog. Nothing is committed; preview with
# `hugo serve`, then run scripts/publish.sh.
#
#   BLOG_OUTBOX=<folder> scripts/outbox.sh
#
# Default outbox: iCloud Drive/Chronicles Blog. A post is matched by its
# `chronicles` id, so a post whose slug or date changed replaces its old
# folder. Images already in content/blog under the same name are kept, since
# scripts/ may have resized them in place (names are content hashes).
set -euo pipefail

cd "$(dirname "$0")/.."
outbox="${BLOG_OUTBOX:-$HOME/Library/Mobile Documents/com~apple~CloudDocs/Chronicles Blog}"
if [[ ! -d "$outbox" ]]; then
  echo "No outbox at $outbox (set BLOG_OUTBOX)." >&2
  exit 1
fi

chronicles_id() {
  sed -n 's/^chronicles: *{"id":"\([^"]*\)".*/\1/p' "$1" | head -1
}

copied=0
while IFS= read -r index; do
  rel="${index#"$outbox"/}"
  bundle="$(dirname "$rel")"   # <date>/<slug>
  id="$(chronicles_id "$index")"
  dest="content/blog/$bundle"

  if [[ -n "$id" ]]; then
    # The same post under another path: its slug or date changed.
    while IFS= read -r other; do
      old="$(dirname "$other")"
      if [[ "$old" != "$dest" && "$(chronicles_id "$other")" == "$id" ]]; then
        echo "move    ${old#content/blog/} -> $bundle"
        rm -rf "$old"
        rmdir "$(dirname "$old")" 2>/dev/null || true
      fi
    done < <(find content/blog -mindepth 3 -maxdepth 3 -name index.md)
  fi

  if [[ -f "$dest/index.md" && -n "$id" && "$(chronicles_id "$dest/index.md")" != "$id" ]]; then
    echo "skip    $bundle: content/blog already has a different post there" >&2
    continue
  fi

  if [[ -f "$dest/index.md" ]] && cmp -s "$index" "$dest/index.md"; then
    continue
  fi

  mkdir -p "$dest"
  cp "$index" "$dest/index.md"
  src_assets="$(dirname "$index")/assets"
  if [[ -d "$src_assets" ]]; then
    mkdir -p "$dest/assets"
    for f in "$src_assets"/*; do
      [[ -e "$dest/assets/$(basename "$f")" ]] || cp "$f" "$dest/assets/"
    done
  fi
  # Images the post no longer uses.
  if [[ -d "$dest/assets" ]]; then
    for f in "$dest/assets"/*; do
      [[ -e "$src_assets/$(basename "$f")" ]] || rm -f "$f"
    done
    rmdir "$dest/assets" 2>/dev/null || true
  fi
  echo "copy    $bundle"
  copied=$((copied + 1))
done < <(find "$outbox" -mindepth 3 -maxdepth 3 -name index.md | sort)

echo "$copied post(s) copied from the outbox."
