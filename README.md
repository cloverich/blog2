
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