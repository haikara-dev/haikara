import Setting from "@/pages/setting";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Setting", () => {
  it("should render correctly", () => {
    expect(Setting).toBeTruthy();
  });
});
