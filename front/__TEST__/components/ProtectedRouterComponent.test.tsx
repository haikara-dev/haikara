import ProtectedRouterComponent from "@/components/ProtectedRouterComponent";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("ProtectedRouterComponent", () => {
  it("renders", () => {
    expect(ProtectedRouterComponent).toBeTruthy();
  });
});
