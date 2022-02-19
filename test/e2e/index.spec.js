const puppeteer = require('puppeteer')

describe('Index page', () => {
  let page
  let browser

  beforeAll(async () => {
    jest.setTimeout(60000)
    browser = await puppeteer.launch({ args: ['--lang=ja'] })
    page = await browser.newPage()
    await page.goto('http://127.0.0.1:3000', {
      waitUntil: 'load',
      timeout: 60000
    })
  })

  afterAll(async () => {
    await page.close()
    await browser.close()
  })

  it('トップページが表示されること', async () => {
    await Promise.all([
      page.waitForNavigation({ waitUntil: ['load', 'networkidle2'] })
    ])

    const text = await page.evaluate(() => document.body.textContent)
    await expect(text).toContain('市内の最新感染動向')
  }, 90000)

  it('「新型コロナウイルス感染症が心配なときに」ページが表示されること', async () => {
    await Promise.all([
      page.waitForNavigation({ waitUntil: ['load', 'networkidle2'] }),
      page.click("a[href='/flow']")
    ])

    const text = await page.evaluate(() => document.body.textContent)
    await expect(text).toContain('新型コロナウイルス感染症が心配なときに')
  }, 90000)
})
