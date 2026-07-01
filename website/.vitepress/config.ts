import { defineConfig } from 'vitepress'
import { withMermaid } from 'vitepress-plugin-mermaid'

const GITHUB = 'https://github.com/scagogogo/cpe-skills'

const enSidebar = {
  '/en/api/': [
    {
      text: 'API Reference',
      items: [
        { text: 'Overview', link: '/en/api/' },
        { text: 'Core Types', link: '/en/api/types' },
        { text: 'Parsing', link: '/en/api/parsing' },
        { text: 'Matching', link: '/en/api/matching' },
        { text: 'Storage', link: '/en/api/storage' },
        { text: 'Dictionary', link: '/en/api/dictionary' },
        { text: 'NVD Integration', link: '/en/api/nvd' },
        { text: 'WFN', link: '/en/api/wfn' },
        { text: 'Validation', link: '/en/api/validation' },
        { text: 'Sets', link: '/en/api/sets' },
        { text: 'Errors', link: '/en/api/errors' }
      ]
    }
  ],
  '/en/guide/': [
    {
      text: 'Guide',
      items: [
        { text: 'Overview', link: '/en/guide/' },
        { text: 'Basic Parsing', link: '/en/guide/basic-parsing' },
        { text: 'CPE Matching', link: '/en/guide/matching' },
        { text: 'WFN Conversion', link: '/en/guide/wfn-conversion' },
        { text: 'Version Comparison', link: '/en/guide/version-comparison' },
        { text: 'Applicability Language', link: '/en/guide/applicability' },
        { text: 'CPE Sets', link: '/en/guide/sets' },
        { text: 'Advanced Matching', link: '/en/guide/advanced-matching' },
        { text: 'Storage', link: '/en/guide/storage' },
        { text: 'NVD Integration', link: '/en/guide/nvd-integration' },
        { text: 'CVE Mapping', link: '/en/guide/cve-mapping' }
      ]
    }
  ]
}

const zhSidebar = {
  '/zh/api/': [
    {
      text: 'API 参考',
      items: [
        { text: '概览', link: '/zh/api/' },
        { text: '核心类型', link: '/zh/api/types' },
        { text: '解析功能', link: '/zh/api/parsing' },
        { text: '匹配算法', link: '/zh/api/matching' },
        { text: '存储接口', link: '/zh/api/storage' },
        { text: '字典管理', link: '/zh/api/dictionary' },
        { text: 'NVD集成', link: '/zh/api/nvd' },
        { text: 'WFN格式', link: '/zh/api/wfn' },
        { text: '验证功能', link: '/zh/api/validation' },
        { text: '集合操作', link: '/zh/api/sets' },
        { text: '错误处理', link: '/zh/api/errors' }
      ]
    }
  ],
  '/zh/guide/': [
    {
      text: '使用指南',
      items: [
        { text: '概览', link: '/zh/guide/' },
        { text: '基础解析', link: '/zh/guide/basic-parsing' },
        { text: 'CPE匹配', link: '/zh/guide/matching' },
        { text: 'WFN转换', link: '/zh/guide/wfn-conversion' },
        { text: '版本比较', link: '/zh/guide/version-comparison' },
        { text: '适用性语言', link: '/zh/guide/applicability' },
        { text: 'CPE集合', link: '/zh/guide/sets' },
        { text: '高级匹配', link: '/zh/guide/advanced-matching' },
        { text: '存储操作', link: '/zh/guide/storage' },
        { text: 'NVD集成', link: '/zh/guide/nvd-integration' },
        { text: 'CVE映射', link: '/zh/guide/cve-mapping' }
      ]
    }
  ]
}

export default withMermaid(
  defineConfig({
    title: 'cpe-skills',
    description: 'A comprehensive CPE (Common Platform Enumeration) toolkit for cybersecurity',
    base: '/cpe-skills/',
    cleanUrls: true,

    head: [
      ['link', { rel: 'icon', type: 'image/svg+xml', href: '/cpe-skills/favicon.svg' }]
    ],

    locales: {
      root: {
        label: 'English',
        lang: 'en',
        themeConfig: {
          nav: [
            { text: 'Home', link: '/en/' },
            { text: 'Guide', link: '/en/guide/' },
            { text: 'API', link: '/en/api/' },
            { text: 'GitHub', link: GITHUB }
          ],
          sidebar: enSidebar
        }
      },
      zh: {
        label: '简体中文',
        lang: 'zh-CN',
        themeConfig: {
          nav: [
            { text: '首页', link: '/zh/' },
            { text: '指南', link: '/zh/guide/' },
            { text: 'API', link: '/zh/api/' },
            { text: 'GitHub', link: GITHUB }
          ],
          sidebar: zhSidebar
        }
      }
    },

    themeConfig: {
      socialLinks: [{ icon: 'github', link: GITHUB }],
      footer: {
        message: 'Released under the MIT License.',
        copyright: 'Copyright © 2024-2026 cpe-skills'
      },
      search: {
        provider: 'local'
      }
    },

    mermaid: {
      theme: 'default'
    }
  })
)
