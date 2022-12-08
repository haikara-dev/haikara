import { act, fireEvent, screen, waitFor } from "@testing-library/react";
import { renderWithProviders } from "../../utils/test-utils";
import EditSiteCategoryFormDialog from "@/components/site-category/EditSiteCategoryFormDialog";
import { rest } from "msw";
import { server } from "../../mocks/server";

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

  it("labelを変更してsumimtすると正しくapiが呼ばれること", async () => {
    const mockFn = jest.fn();
    const BACKEND_ADMIN_API_URL: string =
      process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

    server.use(
      rest.put(
        BACKEND_ADMIN_API_URL + "/site-categories/:id",
        async (req, res, ctx) => {
          const { id } = req.params;

          const { label } = await req.json();

          mockFn({ id, label });

          return res(
            ctx.status(200),
            ctx.json({
              id,
              created_at: "2022-11-18T13:25:23+09:00",
              updated_at: "2022-12-08T12:41:12+09:00",
              label,
              edges: {},
            })
          );
        }
      )
    );
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

    const submitButton = screen.getByRole("button", { name: /save/i });
    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() => {
      expect(mockFn).toHaveBeenCalledTimes(1);
      expect(mockFn).toBeCalledWith({ id: "1", label: "変更後のラベル" });
      expect(handleClose).toBeCalled();
    });
  });
});
