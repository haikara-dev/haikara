import { screen } from "@testing-library/react";
import Footer from "@/components/Footer";
import "@testing-library/jest-dom";
import { renderWithProviders } from "../utils/test-utils";

describe("Footer", () => {
  it("render hooter", () => {
    renderWithProviders(<Footer />);
    const footer = screen.getByText(/haikara/i);
    expect(footer).toBeInTheDocument();
  });
});
