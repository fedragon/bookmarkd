import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  fullyParallel: false,
  retries: 0,
  workers: 3,
  reporter: 'dot',
  use: {
    trace: 'on-first-retry',
    channel: 'chrome',
    headless: true,
  },
});
