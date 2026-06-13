import { defineConfig } from 'vitepress'

// English sidebar configuration
const enSidebar = {
  '/api/': [
    {
      text: 'API Reference',
      items: [
        { text: 'Overview', link: '/api/' },
        { text: 'Core Types', link: '/api/types' },
        { text: 'Parsing', link: '/api/parsing' },
        { text: 'Matching', link: '/api/matching' },
        { text: 'Storage', link: '/api/storage' },
        { text: 'Dictionary', link: '/api/dictionary' },
        { text: 'NVD Integration', link: '/api/nvd' },
        { text: 'WFN', link: '/api/wfn' },
        { text: 'Validation', link: '/api/validation' },
        { text: 'Sets', link: '/api/sets' },
        { text: 'Errors', link: '/api/errors' }
      ]
    }
  ],
  '/examples/': [
    {
      text: 'Examples',
      items: [
        { text: 'Overview', link: '/examples/' },
        { text: 'Basic Parsing', link: '/examples/basic-parsing' },
        { text: 'CPE Matching', link: '/examples/matching' },
        { text: 'WFN Conversion', link: '/examples/wfn-conversion' },
        { text: 'Version Comparison', link: '/examples/version-comparison' },
        { text: 'Applicability Language', link: '/examples/applicability' },
        { text: 'CPE Sets', link: '/examples/sets' },
        { text: 'Advanced Matching', link: '/examples/advanced-matching' },
        { text: 'Storage', link: '/examples/storage' },
        { text: 'NVD Integration', link: '/examples/nvd-integration' },
        { text: 'CVE Mapping', link: '/examples/cve-mapping' }
      ]
    }
  ]
}

// Chinese sidebar configuration
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
  '/zh/examples/': [
    {
      text: '使用示例',
      items: [
        { text: '概览', link: '/zh/examples/' },
        { text: '基础解析', link: '/zh/examples/basic-parsing' },
        { text: 'CPE匹配', link: '/zh/examples/matching' },
        { text: 'WFN转换', link: '/zh/examples/wfn-conversion' },
        { text: '版本比较', link: '/zh/examples/version-comparison' },
        { text: '适用性语言', link: '/zh/examples/applicability' },
        { text: 'CPE集合', link: '/zh/examples/sets' },
        { text: '高级匹配', link: '/zh/examples/advanced-matching' },
        { text: '存储操作', link: '/zh/examples/storage' },
        { text: 'NVD集成', link: '/zh/examples/nvd-integration' },
        { text: 'CVE映射', link: '/zh/examples/cve-mapping' }
      ]
    }
  ]
}

export default defineConfig({
  title: 'CPE Library',
  description: 'Common Platform Enumeration Library for Go',
  base: '/cpe/',
  ignoreDeadLinks: true,

  locales: {
    root: {
      label: 'English',
      lang: 'en',
      title: 'CPE Library',
      description: 'Common Platform Enumeration Library for Go',
      themeConfig: {
        nav: [
          { text: 'Home', link: '/' },
          { text: 'API Reference', link: '/api/' },
          { text: 'Examples', link: '/examples/' },
          { text: 'GitHub', link: 'https://github.com/scagogogo/cpe' }
        ],
        sidebar: enSidebar
      }
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN',
      title: 'CPE 库',
      description: 'Go语言通用平台枚举库',
      themeConfig: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: 'API 参考', link: '/zh/api/' },
          { text: '使用示例', link: '/zh/examples/' },
          { text: 'GitHub', link: 'https://github.com/scagogogo/cpe' }
        ],
        sidebar: zhSidebar
      }
    }
  },

  themeConfig: {
    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/cpe' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024 CPE Library'
    }
  }
})
