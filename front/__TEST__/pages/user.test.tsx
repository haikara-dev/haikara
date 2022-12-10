import Users from "@/pages/user";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Users", () => {
  it("should render", () => {
    expect(Users).toBeTruthy();
  });
});
