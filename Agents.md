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
- **Assets**: 
  - Post-specific assets (like images) are typically placed in an `assets/` subfolder within the post's directory.
  - Global static files (e.g., icons, favicon) in `static/`.
- **Build & Deployment**: 
  - GitHub Actions automates the build and deployment process via `.github/workflows/hugo.yml`.
  - Production files are generated into `public/`.
- **Custom Scripts**:
  - The `scripts/` folder contains Go-based utilities for site maintenance (e.g., image resizing).
- **Git Strategy**:
  - Direct pushes or PRs to the main branch trigger the deployment workflow.
