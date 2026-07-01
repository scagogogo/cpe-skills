import DefaultTheme from 'vitepress/theme'
import type { Theme } from 'vitepress'
import './custom.css'

// Mermaid is wired in via `withMermaid()` in config.ts.
// custom.css overrides the default indigo brand color with a green scale.
export default {
  extends: DefaultTheme
} satisfies Theme
