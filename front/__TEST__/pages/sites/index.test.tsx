import Sites from "@/pages/sites";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Sites", () => {
  it("should render", () => {
    expect(Sites).toBeTruthy();
  });
});
