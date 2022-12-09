import AdminDashBord from "@/components/dashboard/AdminDashBord";
import { server } from "../../mocks/server";
import { rest } from "msw";
import { renderWithProviders } from "../../utils/test-utils";
import { screen } from "@testing-library/react";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

describe("AdminDashBord", () => {
  it("render DashBord", async () => {
    server.use(
      rest.get(BACKEND_ADMIN_API_URL + "/dashboard", (req, res, ctx) => {
        // If authenticated, return a mocked user details
        return res(
          ctx.status(200),
          ctx.json({
            siteSize: 57,
            articleSize: 9609,
            ogpImageSize: 9609,
            feedSize: 4363,
            userSize: 2,
          })
        );
      })
    );

    renderWithProviders(<AdminDashBord />);

    expect(screen.getByText(/Loading/i)).toBeInTheDocument();

    expect(await screen.findByText(/Site: 57/i)).toBeInTheDocument();
    expect(await screen.findByText(/Feed: 4363/i)).toBeInTheDocument();
    expect(await screen.findByText(/Article: 9609/i)).toBeInTheDocument();
    expect(await screen.findByText(/OGP Image: 9609/i)).toBeInTheDocument();
    expect(await screen.findByText(/User: 2/i)).toBeInTheDocument();
  });
});
