import { test, expect } from "@playwright/test";

const fillForm = async (page, { title, description, affectedService }) => {
  await page.getByTitle("create-incident").click();
  await page.getByLabel("Title").fill(title);
  await page.getByLabel("Description").fill(description);
  await page.getByLabel("Affected Service").fill(affectedService);
};

test("submit incident and view in list", async ({ page }) => {
  await page.goto("/");

  // Fill the form
  await fillForm(page, {
    title: "Playwright Incident",
    description: "Network dropped for all users",
    affectedService: "VPN",
  });

  // Submit
  await page.getByRole("button", { name: /submit/i }).click();

  // Wait for reload and check that the new item appears
  await expect(page.locator("table")).toContainText("Playwright Incident");
});

test("basic page loads", async ({ page }) => {
  await page.goto("/");
  await expect(page).toHaveTitle(/incidents/i);
});

test("can view incident detail page", async ({ page }) => {
  await page.goto("/");

  // Submit a new incident
  const incident = {
    title: "VPN Down",
    description: "Users disconnected during work",
    affectedService: "VPN",
  };

  fillForm(page, incident);
  await page.getByRole("button", { name: /submit/i }).click();

  // Wait for it to appear in table
  const row = await page.getByText(incident.title);
  await expect(row).toBeVisible();

  // Click on the title link (to go to /incidents/:id)
  await row.click();

  // Now check details are shown
  await expect(page).toHaveURL(/\/incidents\/\d+/);
  await expect(page.getByTestId("incident-title")).toBeVisible();
  await expect(page.getByTestId("incident-description")).toBeVisible();
  await expect(page.getByTestId("incident-affected-service")).toBeVisible();
  await expect(page.getByTestId("incident-ai-severity")).toBeVisible();
  await expect(page.getByTestId("incident-ai-category")).toBeVisible();
});
