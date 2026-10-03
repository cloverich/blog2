
# Development

Install Go, then:

```sh
# Can pin the version to the version specified in .github/workflows/hugo.yml
go install -tags extended github.com/gohugoio/hugo@latest

# Run locally:
# -D to include drafts
hugo serve
```

Deployment by pushing changes here; this is configured in `.github/workflows/hugo.yml`

## Resizing Images
* TODO: Re-size images in the repo, and pre-commit check to ensure I don't commit large files

An image re-size pipeline is setup in `config.toml`, to constrain output images to 800px in width. Images in blog posts utilize it by adding `?resize=blogImages`. I use this for all images today, but a better solution would be to permanently resize the images I don't care about large sizes on (inline blog posts).

## Post header images

To place an image below navigation and above a post's date/title,
add a bundle-relative image and alt text to its front matter:

```yaml
headerImage:
  src: assets/beach.jpeg
  alt: A southern Oregon beach
```

Header images use an 8:3 panorama crop, reach 720px wide, and fit within the
viewport. The title is left aligned, with metadata below it. Keep screenshots and
other images that need to be read in the Markdown body.

The site follows the system color scheme automatically through CSS. Code uses
one muted hue per theme, with no syntax highlighting or separate background.


## Local header studies

Run `hugo server -D` and open `/header-lab/` for eight header compositions
using the Chronicles post. The gallery has full-post previews and a theme
selector. Study pages are drafts and excluded from normal production builds.


Blog metadata labels every post “Human written” and includes the first eight
hex characters of SHA-256 of the raw Markdown body. The hash is deterministic
and changes when the body changes; front matter and theme changes do not affect
it. Hover the hash for its full value, or copy its link as a version marker.
The marker identifies content, not an archived copy of that version.
