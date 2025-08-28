import { test, expect } from '@playwright/test'

test('convert action should not return HTTP 400 for single-part volume mappings', async ({ page }) => {
  // Navigate to the app root (playwright config baseURL)
  await page.goto('/')

  // Example docker-compose with a single-part volume mapping that used to trigger a 400
  const compose = `version: '3.8'
services:
  api:
    image: node:18
    volumes:
      - /app/node_modules
    ports:
      - "3000:3000"
`

  // Fill the compose textarea (assumes textarea has id or placeholder we can target)
  const textarea = await page.locator('textarea, [data-test="compose-input"]').first()
  await textarea.fill(compose)

  // Intercept the conversion request
  const [response] = await Promise.all([
    page.waitForResponse(resp => resp.url().includes('/api/v1/convert/') && resp.request().method() === 'POST'),
    // Click the convert button (assumes a button with text 'Convertir' or 'Convert')
    page.locator('button:has-text("Convert") , button:has-text("Convertir")').first().click(),
  ])

  // Assert we did not receive HTTP 400
  expect(response.status(), 'convert API should not return 400').not.toBe(400)

  // Optionally parse JSON and log errors/warnings for debugging
  const body = await response.json().catch(() => null)
  if (body && body.errors && body.errors.length) {
    // Ensure errors don't include the specific parse error code (optional)
    expect(body.success === false ? body.errors.length : 0).toBeLessThan(2)
  }
})
