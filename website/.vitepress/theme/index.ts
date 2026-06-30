import DefaultTheme from 'vitepress/theme'
import type { Theme } from 'vitepress'

// Mermaid is wired in via `withMermaid()` in config.ts; no extra theme work needed.
export default {
  extends: DefaultTheme
} satisfies Theme
