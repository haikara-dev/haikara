import SiteEdit from "@/pages/sites/[id]/edit";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("SiteEdit", () => {
  it("should render", () => {
    expect(SiteEdit).toBeTruthy();
  });
});
