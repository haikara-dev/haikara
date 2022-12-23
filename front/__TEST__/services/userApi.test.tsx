import { useGetArticlesQuery } from "@/services/userApi";

describe("userApi", () => {
  it("renders", () => {
    expect(useGetArticlesQuery).toBeTruthy();
  });
});
