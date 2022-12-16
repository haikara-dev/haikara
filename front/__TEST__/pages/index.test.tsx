import { screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Home from "@/pages/index";
import { mockAuthorizedAuth, renderWithProviders } from "../utils/test-utils";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("Home", () => {
  it("render heading", () => {
    renderWithProviders(<Home />, {
      router: {},
    });
    const heading = screen.getByText(/haikara/i);
    expect(heading).toBeInTheDocument();
  });
  it("render heading with Login(role user)", async () => {
    renderWithProviders(<Home />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("user") },
    });

    const heading = screen.getByText(/haikara/i);
    expect(heading).toBeInTheDocument();

    const loading = screen.getByText(/Loading/i);
    expect(loading).toBeInTheDocument();

    const articles = await screen.findByText(/記事タイトル1/i);
    expect(articles).toBeInTheDocument();
  });
});
