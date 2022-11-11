import { render, screen } from "@testing-library/react";
import Footer from "@/components/Footer";
import "@testing-library/jest-dom";

describe("Footer", () => {
  it("render hooter", () => {
    render(<Footer />);
    const footer = screen.getByText(/haikara/i);
    expect(footer).toBeInTheDocument();
  });
});
