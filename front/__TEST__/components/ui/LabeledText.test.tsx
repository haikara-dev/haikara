import { screen } from "@testing-library/react";
import LabeledText from "@/components/ui/LabeledText";
import { render } from "@testing-library/react";

describe("LabeledText", () => {
  it("renders", () => {
    render(<LabeledText label="ラベル" value="テスト" />);
    expect(screen.getByRole("label", { name: "ラベル" })).toBeInTheDocument();
    expect(screen.getByRole("textbox")).toHaveValue("テスト");
  });
});
