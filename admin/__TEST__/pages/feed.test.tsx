import Feeds from "@/pages/feed";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Feed", () => {
  it("should render correctly", () => {
    expect(Feeds).toBeTruthy();
  });
});
