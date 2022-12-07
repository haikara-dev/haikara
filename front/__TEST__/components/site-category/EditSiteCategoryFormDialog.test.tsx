import { act, fireEvent, screen, waitFor } from "@testing-library/react";
import { renderWithProviders } from "../../utils/test-utils";
import EditSiteCategoryFormDialog from "@/components/site-category/EditSiteCategoryFormDialog";

describe("EditSiteCategoryFormDialog", () => {
  it("inputのデフォルト値にsiteCategoryのlabelが設定されている", () => {
    const handleClose = jest.fn();
    const siteCategory = {
      id: 1,
      label: "ラベル",
    };
    renderWithProviders(
      <EditSiteCategoryFormDialog
        open={true}
        handleClose={handleClose}
        siteCategory={siteCategory}
      />
    );

    const input = screen.getByRole("textbox", {
      name: "Site Category Label",
    }) as HTMLInputElement;

    expect(input.value).toEqual(siteCategory.label);
  });

  it("labelを入力できる", () => {
    const handleClose = jest.fn();
    const siteCategory = {
      id: 1,
      label: "ラベル",
    };
    renderWithProviders(
      <EditSiteCategoryFormDialog
        open={true}
        handleClose={handleClose}
        siteCategory={siteCategory}
      />
    );

    const input = screen.getByRole("textbox", {
      name: "Site Category Label",
    }) as HTMLInputElement;

    fireEvent.change(input, { target: { value: "変更後のラベル" } });

    expect(input.value).toEqual("変更後のラベル");
  });

  it.todo("labelを変更してsumimtすると正しくapiが呼ばれること");
});
