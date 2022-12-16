import { render, screen } from "@testing-library/react";
import PaginationHeader from "@/components/ui/PaginationHeader";

describe("PaginationHeader", () => {
  it("renders", () => {
    render(
      <PaginationHeader
        totalCount={996}
        page={10}
        pageSize={100}
        dataSize={96}
      />
    );
    expect(screen.getByText("996件中 901 - 996件")).toBeInTheDocument();
  });
});
