import Register from "@/pages/register";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Register", () => {
  it("should render correctly", () => {
    expect(Register).toBeTruthy();
  });
});
