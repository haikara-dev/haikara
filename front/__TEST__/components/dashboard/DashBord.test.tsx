import { screen } from "@testing-library/react";
import DashBord from "@/components/dashboard/DashBord";
import "@testing-library/jest-dom";
import { renderWithProviders } from "../../utils/test-utils";
import { server } from "../../mocks/server";
import { rest } from "msw";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

describe("DashBord", () => {
  it("render DashBord", async () => {
    server.use(
      rest.get(BACKEND_API_URL + "/dashboard", (req, res, ctx) => {
        // If authenticated, return a mocked user details
        return res(
          ctx.status(200),
          ctx.json({
            siteSize: 57,
            articleSize: 6810,
          })
        );
      })
    );

    renderWithProviders(<DashBord />);

    expect(screen.getByText(/Loading/i)).toBeInTheDocument();

    expect(await screen.findByText(/Site: 57/i)).toBeInTheDocument();
    expect(await screen.findByText(/Article: 6810/i)).toBeInTheDocument();
  });
});
