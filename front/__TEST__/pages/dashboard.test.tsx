import DashboardPage from "@/pages/dashboard";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("DashboardPage", () => {
  it("should render", () => {
    expect(DashboardPage).toBeTruthy();
  });
});
