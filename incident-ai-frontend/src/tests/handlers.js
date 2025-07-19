import { http } from "msw";
import { BASE_URL } from "../api/incident";

export const handlers = [
  http.get(BASE_URL, (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json([
        {
          id: 1,
          title: "Net down",
          ai_category: "Network",
          ai_severity: "High",
          affected_service: "Router",
          created_at: "2025-07-19",
        },
      ])
    );
  }),
];
