import Site from "@/pages/sites/[id]";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Site", () => {
  it("should render", () => {
    expect(Site).toBeTruthy();
  });
});
