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

To change it:

1. Open the lab (https://cloverich.github.io/artifacts/, Horizon Plate), drop in
   the source, and paste `recipe.json` onto the page to restore these settings.
2. Tune, then **Copy recipe** and replace `recipe.json` with it.
3. Re-render all four files from a checkout of cloverich/artifacts:

   ```sh
   node horizon-plate/build.mjs
   node horizon-plate/render.mjs <source> ../blog2/static/footer/recipe.json ../blog2/static/footer --dark
   ```

   (Or hand an agent the recipe and the source image; never exported PNGs
   through chat, which re-encodes them.)
4. If you changed a plate's width, height, or Position X, update the matching
   values in `.site-footer` in `static/new.css` (Copy CSS in the lab shows them).
