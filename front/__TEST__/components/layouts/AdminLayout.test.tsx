import AdminLayout from "@/components/layouts/AdminLayout";
import { renderWithProviders } from "../../utils/test-utils";
import { screen } from "@testing-library/react";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("AdminLayout", () => {
  it("should render", () => {
    renderWithProviders(<AdminLayout>test</AdminLayout>);
    expect(screen.getByText("test")).toBeInTheDocument();
  });
  it.todo("setOpen");
});
