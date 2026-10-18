# Footer image

The post footer's landscape is four alpha masks made in the Horizon Plate lab
(cloverich/artifacts, `horizon-plate/`). CSS paints the ink colour through them
(`mask-image`), so the PNGs carry coverage only.

- `plate-wide.png`, `plate-narrow.png`: light mode, from `recipe.json` as is.
- `plate-wide-dark.png`, `plate-narrow-dark.png`: dark mode, the same recipe with
  `"inkLight": true`, so light ink paints the light areas (a print, not a negative).
- The narrow pair is used below 820 px, the wide pair from there up. Each visitor
  downloads one file.

Source: `watched_and_acceptance.png`, a dream of the author's, drawn by them and
rendered as a painting by Gemini. The source is not committed.

To change it: load the source in the lab, paste `recipe.json` onto the page, tune,
Copy recipe, and re-export (or have an agent render the four files headless from
the recipe and source). Plate geometry lives in the `.site-footer` rules in
`static/new.css`; keep their sizes and positions in sync with the recipe.
