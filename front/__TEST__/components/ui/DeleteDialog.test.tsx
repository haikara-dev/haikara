import { getNodeText, screen } from "@testing-library/react";
import DeleteDialog from "@/components/ui/DeleteDialog";
import { render } from "@testing-library/react";

describe("DeleteDialog", () => {
  it("should render", () => {
    render(
      <DeleteDialog
        open={true}
        title="記事の削除"
        cancelHandler={jest.fn()}
        agreeHandler={jest.fn()}
      >
        削除しますか？
      </DeleteDialog>
    );

    expect(
      screen.getByRole("heading", {
        name: "記事の削除",
      })
    ).toBeInTheDocument();

    expect(screen.getByText("削除しますか？")).toBeInTheDocument();

    const cancelButton = screen.getByTestId("cancel-button");
    const agreeButton = screen.getByTestId("agree-button");

    expect(cancelButton).toBeInTheDocument();
    expect(agreeButton).toBeInTheDocument();

    expect(getNodeText(cancelButton)).toBe("Chancel");
    expect(getNodeText(agreeButton)).toBe("Agree");
  });

  it("should render label Chancel button", () => {
    render(
      <DeleteDialog
        open={true}
        title="記事の削除"
        cancelHandler={jest.fn()}
        agreeHandler={jest.fn()}
        cancelButtonLabel="キャンセル"
      >
        削除しますか？
      </DeleteDialog>
    );

    const cancelButton = screen.getByTestId("cancel-button");
    expect(getNodeText(cancelButton)).toBe("キャンセル");
  });

  it("should render label Agree button", () => {
    render(
      <DeleteDialog
        open={true}
        title="記事の削除"
        cancelHandler={jest.fn()}
        agreeHandler={jest.fn()}
        agreeButtonLabel="削除"
      >
        削除しますか？
      </DeleteDialog>
    );

    const agreeButton = screen.getByTestId("agree-button");
    expect(getNodeText(agreeButton)).toBe("削除");
  });

  it.todo("open, close");
});
