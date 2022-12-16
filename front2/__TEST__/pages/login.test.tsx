import Login from "@/pages/login";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Login", () => {
  it("should render correctly", () => {
    expect(Login).toBeTruthy();
  });
});
