(() => {
  const root = document.documentElement;
  if (new URLSearchParams(location.search).has('embed')) root.classList.add('lab-embedded');
  const select = document.querySelector('#lab-theme');
  function apply(theme) {
    if (theme === 'system') delete root.dataset.labTheme;
    else root.dataset.labTheme = theme;
    if (select) select.value = theme;
  }
  function savedTheme() {
    try { return localStorage.getItem('header-lab-theme') || 'system'; } catch { return 'system'; }
  }
  const requested = new URLSearchParams(location.search).get('theme');
  apply(['system','light','dark'].includes(requested) ? requested : savedTheme());
  function syncPreviews(theme) {
    document.querySelectorAll('iframe').forEach(frame => {
      const url = new URL(frame.src);
      url.searchParams.set('theme', theme);
      frame.src = url.href;
    });
  }
  syncPreviews(select?.value || savedTheme());
  select?.addEventListener('change', () => {
    apply(select.value);
    try { localStorage.setItem('header-lab-theme', select.value); } catch {}
    syncPreviews(select.value);
  });
  addEventListener('message', event => {
    if (event.origin === location.origin && ['system','light','dark'].includes(event.data?.headerLabTheme)) apply(event.data.headerLabTheme);
  });
  addEventListener('storage', event => { if (event.key === 'header-lab-theme') apply(savedTheme()); });
  const thumbnails = document.querySelectorAll('.lab-thumbnail');
  const observer = new ResizeObserver(entries => entries.forEach(({target,contentRect}) => {
    const frame = target.querySelector('iframe');
    frame.style.transform = `scale(${contentRect.width / 1280})`;
    target.style.height = `${Math.round(contentRect.width / 1280 * 850)}px`;
  }));
  thumbnails.forEach(el => observer.observe(el));
})();
