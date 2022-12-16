import AdminLayout from "@/components/layouts/AdminLayout";
import { renderWithProviders } from "../../utils/test-utils";
import { fireEvent, screen } from "@testing-library/react";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("AdminLayout", () => {
  it("should render children", () => {
    renderWithProviders(<AdminLayout>test</AdminLayout>);
    expect(screen.getByText("test")).toBeInTheDocument();
  });

  it("setOpen", () => {
    renderWithProviders(<AdminLayout>test</AdminLayout>);

    const button = screen.getByTestId("MenuIcon");
    const main = screen.getByTestId("AdminLayout-Main");

    // open true
    expect(main).toHaveStyle("margin-left: 0");
    fireEvent.click(button);
    // open false
    expect(main).toHaveStyle("margin-left: -140px");
  });
});
