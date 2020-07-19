export default {
  strategy: 'prefix_except_default',
  detectBrowserLanguage: {
    useCookie: true,
    cookieKey: 'i18n_redirected'
  },
  defaultLocale: 'ja',
  vueI18n: {
    fallbackLocale: 'ja',
    formatFallbackMessages: true
  },
  // vueI18nLoader: true,
  lazy: true,
  langDir: 'assets/locales/',
  locales: [
    {
      code: 'ja',
      name: '日本語',
      iso: 'ja-JP',
      file: 'ja.json',
      description: 'Japanese'
    },
    {
      code: 'en',
      name: 'English',
      iso: 'en-US',
      file: 'en.json',
      description: 'English'
    },
    {
      code: 'pt-br',
      name: 'Portuguese',
      iso: 'pt-BR',
      file: 'pt_BR.json',
      description: 'Portuguese'
    },
    {
      code: 'ja-basic',
      name: 'やさしい にほんご',
      iso: 'ja-JP',
      file: 'ja-Hira.json',
      description: 'Easy Japanese'
    },
    {
      code: 'tl-ph',
      name: 'タガログ語',
      iso: 'tl',
      file: 'tl.json',
      description: 'Tagalog'
    },
    {
      code: 'zh-cn',
      name: '简体中文',
      iso: 'zh-CN',
      file: 'zh_CN.json',
      description: 'Simplified Chinese'
    },
    {
      code: 'vi',
      name: 'Tiếng Việt',
      iso: 'vi',
      file: 'vi.json',
      description: 'Tiếng Việt'
    }
  ]
}
