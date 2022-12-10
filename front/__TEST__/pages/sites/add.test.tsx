import SiteAdd from "@/pages/sites/add";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("SiteAdd", () => {
  it("should render", () => {
    expect(SiteAdd).toBeTruthy();
  });
});
