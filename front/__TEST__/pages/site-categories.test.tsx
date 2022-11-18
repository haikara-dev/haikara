import { screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { mockAuthorizedAuth, renderWithProviders } from "../utils/test-utils";
import SiteCategories from "@/pages/site-categories";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("SiteCategories", () => {
  it("render heading", () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
    });
    const heading = screen.getByText(/Site Category/i);
    expect(heading).toBeInTheDocument();
  });
  it("render heading with Login(role user)", async () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });

    const heading = screen.getByText(/Site Category/i);
    expect(heading).toBeInTheDocument();

    const loading = screen.getByText(/Loading/i);
    expect(loading).toBeInTheDocument();

    const articles = await screen.findByText(/メディア/i);
    expect(articles).toBeInTheDocument();
  });
});
