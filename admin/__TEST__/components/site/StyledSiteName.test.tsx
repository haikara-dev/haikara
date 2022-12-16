import { screen } from "@testing-library/react";
import { renderWithProviders } from "../../utils/test-utils";
import StyledSiteName from "@/components/site/StyledSiteName";
describe("StyledSiteName", () => {
  it("サイト名が分離される", () => {
    renderWithProviders(<StyledSiteName>test/fashion</StyledSiteName>);
    expect(screen.queryByText("test/fashion")).not.toBeInTheDocument();
    expect(screen.getByText("test")).toBeInTheDocument();
    expect(screen.getByText("/fashion")).toBeInTheDocument();
  });
});
