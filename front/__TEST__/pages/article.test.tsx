import Articles from "@/pages/article";
jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});
describe("Article", () => {
  it("should", () => {
    expect(Articles).toBeTruthy();
  });
});
