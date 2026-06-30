---
layout: page
---

<script setup>
// Redirect root to the English home page.
// root locale maps to English content under /en/.
if (typeof window !== 'undefined') {
  const base = import.meta.env.BASE_URL // "/cpe-skills/"
  window.location.replace(base + 'en/')
}
</script>

<meta http-equiv="refresh" content="0; url=./en/" />

Redirecting to the English home page… If you are not redirected automatically, follow [this link](./en/).
