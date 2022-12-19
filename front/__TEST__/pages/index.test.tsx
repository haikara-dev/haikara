import { screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Home from "@/pages/index";
import { renderWithProviders } from "../utils/test-utils";

describe("Home", () => {
  it("render heading", () => {
    renderWithProviders(<Home />, {
      router: {},
    });
    const heading = screen.getByText(/haikara/i);
    expect(heading).toBeInTheDocument();
  });
});
