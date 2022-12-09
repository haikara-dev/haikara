import AuthLayout from "@/components/layouts/AuthLayout";
import { renderWithProviders } from "../../utils/test-utils";
import { screen } from "@testing-library/react";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("AuthLayout", () => {
  it("should render children", () => {
    renderWithProviders(<AuthLayout>test</AuthLayout>);
    expect(screen.getByText("test")).toBeInTheDocument();
  });
});
