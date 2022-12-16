import { useAuth } from "@/services/firebase";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("useAuth", () => {
  it("renders", () => {
    expect(useAuth).toBeTruthy();
  });
});
