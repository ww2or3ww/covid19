import { Configuration } from '@nuxt/types'
import { Configuration as WebpackConfiguration } from 'webpack'
import i18n from './nuxt-i18n.config'
const webpack = require('webpack')
const purgecss = require('@fullhuman/postcss-purgecss')
const autoprefixer = require('autoprefixer')
const environment = process.env.NODE_ENV || 'development'

const config: Configuration = {
  /*
   ** Headers of the page
   */
  head: {
    htmlAttrs: {
      prefix: 'og: http://ogp.me/ns#'
    },
    titleTemplate: '%s | 浜松市 新型コロナウイルス感染症対策サイト',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'og:type', property: 'og:type', content: 'website' },
      {
        hid: 'og:url',
        property: 'og:url',
        content: 'https://https://stopcovid19.code4hamamatsu.org/'
      },
      {
        hid: 'twitter:card',
        name: 'twitter:card',
        content: 'summary_large_image'
      },
      {
        hid: 'twitter:site',
        name: 'twitter:site',
        content: '@Hamamatsu_pr'
      },
      {
        hid: 'twitter:creator',
        name: 'twitter:creator',
        content: '@Hamamatsu_pr'
      },
      // {
      //   hid: 'fb:app_id',
      //   property: 'fb:app_id',
      //   content: '2879625188795443'
      // },
      {
        hid: 'note:card',
        property: 'note:card',
        content: 'summary_large_image'
      }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/hamamatsu/favicon.ico' },
      {
        rel: 'shortcut icon',
        type: 'image/x-icon',
        href: '/hamamatsu/favicon.ico'
      },
      { rel: 'apple-touch-icon', href: '/hamamatsu/apple-touch-icon.png' },
      {
        rel: 'stylesheet',
        href: 'https://use.fontawesome.com/releases/v5.6.1/css/all.css'
      }
    ]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: ['~assets/global.scss'],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    {
      src: '@/plugins/vue-chart.ts',
      ssr: true
    },
    {
      src: '@/plugins/axe',
      ssr: true
    },
    {
      src: '@/plugins/vuetify.ts',
      ssr: true
    }
  ],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    '@nuxtjs/stylelint-module',
    '@nuxtjs/vuetify',
    '@nuxt/typescript-build',
    '@nuxtjs/google-analytics'
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    '@nuxtjs/pwa',
    // Doc: https://github.com/nuxt-community/dotenv-module
    ['@nuxtjs/dotenv', { filename: `.env.${environment}` }],
    ['nuxt-i18n', i18n],
    'nuxt-svg-loader',
    'nuxt-purgecss',
    ['vue-scrollto/nuxt', { duration: 1000, offset: -72 }]
  ],
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    defaultAssets: {
      icons: false
    }
  },
  googleAnalytics: {
    id: 'UA-9971993-6'
  },
  build: {
    plugins: [
      new webpack.ProvidePlugin({
        mapboxgl: 'mapbox-gl'
      })
    ],
    postcss: {
      plugins: [
        autoprefixer({ grid: 'autoplace' }),
        purgecss({
          content: [
            './pages/**/*.vue',
            './layouts/**/*.vue',
            './components/**/*.vue',
            './node_modules/vuetify/dist/vuetify.js',
            './node_modules/vue-spinner/src/ScaleLoader.vue'
          ],
          whitelist: ['html', 'body', 'nuxt-progress', 'DataCard'],
          whitelistPatterns: [/(col|row)/]
        })
      ]
    },
    extend(config: WebpackConfiguration) {
      // default externals option is undefined
      config.externals = [{ moment: 'moment' }]
    }
    // https://ja.nuxtjs.org/api/configuration-build/#hardsource
    // hardSource: process.env.NODE_ENV === 'development'
  },
  manifest: {
    name: '浜松市 新型コロナウイルス感染症対策サイト',
    description: '浜松市 新型コロナウイルス感染症対策サイト',
    theme_color: '#00a040',
    background_color: '#ffffff',
    display: 'standalone',
    Scope: '/',
    start_url: '/',
    splash_pages: null
  },
  pwa: {
    icon: {
      iconSrc: 'static/hamamatsu/icon.png'
    },
    workbox: {
      enabled: false
    }
  },
  generate: {
    fallback: true,
    routes() {
      const locales = [
        'ja',
        'ja-basic',
        'en',
        'pt-br',
        'tl',
        'zh-cn',
        'vi',
        'es'
      ]
      const pages = [
        '/cards/details-of-confirmed-cases',
        '/cards/number-of-confirmed-cases',
        '/cards/attributes-of-confirmed-cases',
        '/cards/number-of-inspection-persons',
        '/cards/number-of-reports-to-covid19-telephone-advisory-center'
      ]

      const routes: string[] = []
      locales.forEach(locale => {
        pages.forEach(page => {
          if (locale === 'ja') {
            routes.push(page)
            return
          }
          const route = `/${locale}${page}`
          routes.push(route)
        })
      })
      return routes
    }
  },
  // /*
  // ** hot read configuration for docker
  // */
  watchers: {
    webpack: {
      poll: true
    }
  }
}

export default config
