import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: "./src/tests",
  testMatch: "**/*.spec.{js,jsx,ts,tsx}",
  use: {
    baseURL: "http://localhost:5173",
  },
  webServer: {
    command: "npm run dev",
    port: 5173,
    timeout: 60 * 1000,
    reuseExistingServer: true,
  },
});
