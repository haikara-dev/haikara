import { screen } from "@testing-library/react";
import DrawerMenu from "@/components/DrawerMenu";
import { mockAuthorizedAuth, renderWithProviders } from "../utils/test-utils";

describe("DrawerMenu", () => {
  it("render menu for guest", async () => {
    renderWithProviders(<DrawerMenu open={true} />, {
      router: {},
    });
    expect(
      screen.queryByRole("link", { name: "Dashboard" })
    ).not.toBeInTheDocument();
    expect(
      screen.queryByRole("link", { name: "Article" })
    ).not.toBeInTheDocument();
  });
  it("render menu for user", async () => {
    renderWithProviders(<DrawerMenu open={true} />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("user") },
    });
    expect(screen.getByRole("link", { name: "Dashboard" })).toBeInTheDocument();
    expect(
      screen.queryByRole("link", { name: "Article" })
    ).not.toBeInTheDocument();
  });
  it("render menu for admin", async () => {
    renderWithProviders(<DrawerMenu open={true} />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });
    expect(screen.getByRole("link", { name: "Dashboard" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Article" })).toBeInTheDocument();
  });
});
