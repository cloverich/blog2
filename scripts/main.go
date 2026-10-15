// Command images shrinks blog images to the size they are displayed at.
//
// Header images (front matter headerImage.src) are capped at heroWidth and
// all other images at bodyWidth: twice the CSS width they render at, for
// high-density screens. An image is only rewritten when that saves at least
// minSaving of its size on disk, so running it again is a no-op.
//
//	go -C scripts run .          # report what would change
//	go -C scripts run . -write   # rewrite images in place
//	go -C scripts run . -check   # exit 1 if anything would change
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	heroWidth   = 1600 // .post-hero renders at up to 720px
	bodyWidth   = 1200 // body text column is 600px
	jpegQuality = 80
	// JPEGs at or under the width limit are re-encoded only above this size.
	jpegMaxBytes = 600 * 1024
	// Rewrites must save at least this fraction of the file, to avoid churn.
	minSaving = 0.15
)

var (
	writeFlag = flag.Bool("write", false, "rewrite images in place")
	checkFlag = flag.Bool("check", false, "exit 1 if any image would be rewritten")
	dirFlag   = flag.String("dir", "", "directory to scan (default: the site's content/)")
)

var headerSrc = regexp.MustCompile(`(?m)^headerImage:\s*\n(?:[ \t]+.*\n)*?[ \t]+src:\s*["']?([^"'\n]+?)["']?\s*$`)

func main() {
	flag.Parse()

	dir := *dirFlag
	if dir == "" {
		root, err := siteRoot()
		if err != nil {
			fail(err)
		}
		dir = filepath.Join(root, "content")
	}

	pending := 0
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || formatOf(path) < 0 {
			return err
		}
		changed, err := process(path)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if changed {
			pending++
		}
		return nil
	})
	if err != nil {
		fail(err)
	}

	if *checkFlag && pending > 0 {
		fmt.Printf("%d image(s) need resizing; run: go -C scripts run . -write\n", pending)
		os.Exit(1)
	}
}

// process reports (and with -write, applies) the change for one image.
func process(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	src, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return false, err
	}

	limit := bodyWidth
	if isHeaderImage(path) {
		limit = heroWidth
	}
	format := imaging.Format(formatOf(path))
	w := src.Bounds().Dx()

	tooWide := w > limit
	tooHeavy := format == imaging.JPEG && info.Size() > jpegMaxBytes
	if !tooWide && !tooHeavy {
		return false, nil
	}

	var dst image.Image = src
	if tooWide {
		dst = imaging.Resize(src, limit, 0, imaging.Lanczos)
	}
	var buf bytes.Buffer
	if err := imaging.Encode(&buf, dst, format, imaging.JPEGQuality(jpegQuality)); err != nil {
		return false, err
	}

	rel := relPath(path)
	if float64(buf.Len()) > float64(info.Size())*(1-minSaving) {
		return false, nil
	}

	fmt.Printf("resize  %s  %dpx %s -> %dpx %s\n", rel, w, kb(info.Size()), dst.Bounds().Dx(), kb(int64(buf.Len())))
	if *writeFlag {
		if err := os.WriteFile(path, buf.Bytes(), info.Mode().Perm()); err != nil {
			return false, err
		}
	}
	return true, nil
}

// isHeaderImage reports whether the bundle's index.md uses path as its headerImage.
func isHeaderImage(path string) bool {
	dir := filepath.Dir(path)
	for i := 0; i < 2; i++ { // the image may sit in the bundle's assets/ folder
		md, err := os.ReadFile(filepath.Join(dir, "index.md"))
		if err == nil {
			m := headerSrc.FindSubmatch(md)
			if m == nil {
				return false
			}
			want := filepath.Join(dir, strings.TrimPrefix(string(m[1]), "./"))
			return filepath.Clean(want) == filepath.Clean(path)
		}
		dir = filepath.Dir(dir)
	}
	return false
}

func formatOf(path string) int {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".jpg", ".jpeg":
		return int(imaging.JPEG)
	case ".png":
		return int(imaging.PNG)
	}
	return -1
}

// siteRoot walks up from the working directory to the folder holding config.toml.
func siteRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "config.toml")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("config.toml not found above the working directory")
		}
		dir = parent
	}
}

func relPath(path string) string {
	if root, err := siteRoot(); err == nil {
		if rel, err := filepath.Rel(root, path); err == nil {
			return rel
		}
	}
	return path
}

func kb(n int64) string { return fmt.Sprintf("%dKB", n/1024) }

func fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
