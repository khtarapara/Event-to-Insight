import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import IncidentForm from "./IncidentForm";
import { createIncident } from "../../api/incident";
import { describe, expect, it, vi } from "vitest";

vi.mock("../../api/incident", () => ({
  createIncident: vi.fn(),
}));

describe("IncidentForm", () => {
  it("validates required fields", async () => {
    render(<IncidentForm />);
    fireEvent.click(screen.getByRole("button", { name: /submit/i }));

    expect(await screen.findAllByText("Required")).toHaveLength(3);
  });

  it("submits form data successfully", async () => {
    createIncident.mockResolvedValue({});

    render(<IncidentForm />);
    fireEvent.change(screen.getByLabelText(/title/i), {
      target: { value: "Test" },
    });
    fireEvent.change(screen.getByLabelText(/description/i), {
      target: { value: "Details" },
    });
    fireEvent.change(screen.getByLabelText(/affected service/i), {
      target: { value: "DB" },
    });

    fireEvent.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() => {
      expect(createIncident).toHaveBeenCalledWith({
        title: "Test",
        description: "Details",
        affected_service: "DB",
      });
    });
  });
});
