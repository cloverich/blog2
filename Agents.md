# Repository Overview

This is a personal static site built using [Hugo](https://gohugo.io/).

- **Framework**: Hugo (Static Site Generator)
- **Primary Configuration**: `config.toml`
- **Content**: 
  - Markdown files located in the `content/` directory.
  - Blog posts follow the pattern: `content/blog/YYYY-MM-DD/[slug]/index.md`.
  - Blog posts are structured as leaf bundles (folders with an `index.md`).
- **Archetypes**: 
  - `archetypes/default.md` provides the default front matter template for new content.
- **Layouts & Styling**:
  - Custom HTML templates in `layouts/`.
  - Global CSS files in `static/` (e.g., `new.css`, `syntax.css`, `typography_blog1.css`).
  - Post footer (`layouts/partials/footer.html`, `.site-footer` in `new.css`): on blog posts a landscape rises behind it, four alpha masks in `static/footer/` made in the Horizon Plate lab (cloverich/artifacts). See `static/footer/README.md` to change the image.
- **Assets**: 
  - Post-specific assets (like images) are typically placed in an `assets/` subfolder within the post's directory.
  - Global static files (e.g., icons, favicon) in `static/`.
- **Build & Deployment**: 
  - GitHub Actions automates the build and deployment process via `.github/workflows/hugo.yml`.
  - Production files are generated into `public/`.
- **Custom Scripts**:
  - `scripts/` is a Go program that shrinks post images to their display size (header images 1600px wide, body images 1200px). Run `go -C scripts run . -write` after adding images (no flag = dry run, `-check` exits 1 if anything would change). Commit the resized images, not camera/screenshot originals.
  - `scripts/publish.sh` puts posts that Chronicles wrote into `content/blog` live: resizes images, commits, pushes `master`. `--dry-run` only lists them. Chronicles posts carry `chronicles: {id, revision}` in front matter; that is how Chronicles knows a post is published and whether the note changed since (spec: `chronicles/spec/publish.md`). Hand edits to those posts are overwritten by the next publish from Chronicles.
  - `scripts/outbox.sh` copies posts the iPhone app wrote to its iCloud Drive blog folder (default `iCloud Drive/Chronicles Blog`, or `BLOG_OUTBOX`) into `content/blog`, without committing; preview with `hugo serve`. `publish.sh` runs it first when that folder exists.
- **Git Strategy**:
  - Direct pushes or PRs to the `master` branch trigger the deployment workflow (there is no `main`).
