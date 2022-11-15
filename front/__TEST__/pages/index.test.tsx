import { screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Home from "@/pages/index";
import { renderWithProviders } from "../utils/test-utils";

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
    screen.debug();
    const heading = screen.getByText(/haikara/i);
    expect(heading).toBeInTheDocument();
  });
});
