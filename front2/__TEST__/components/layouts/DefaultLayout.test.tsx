import DefaultLayout from "@/components/layouts/DefaultLayout";
import { renderWithProviders } from "../../utils/test-utils";
import { screen } from "@testing-library/react";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("DefaultLayout", () => {
  it("should render children", () => {
    renderWithProviders(<DefaultLayout>test</DefaultLayout>);
    expect(screen.getByText("test")).toBeInTheDocument();
  });
});
