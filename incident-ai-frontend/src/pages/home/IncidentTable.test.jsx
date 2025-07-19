import { render, screen } from "@testing-library/react";
import IncidentTable from "./IncidentTable";
import { MemoryRouter } from "react-router-dom";
import { expect, test } from "vitest";

test("renders rows from props", () => {
  const sampleData = [
    {
      id: 1,
      title: "VPN Down",
      ai_severity: "High",
      ai_category: "Network",
      affected_service: "VPN",
      created_at: "2025-07-19",
    },
  ];

  render(
    <MemoryRouter>
      <IncidentTable data={sampleData} />
    </MemoryRouter>
  );

  expect(screen.getByText("VPN Down")).toBeInTheDocument();
  expect(screen.getByText("High")).toBeInTheDocument();
  expect(screen.getByText("Network")).toBeInTheDocument();
});
