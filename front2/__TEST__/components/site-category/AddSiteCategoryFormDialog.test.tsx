import { act, fireEvent, screen, waitFor } from "@testing-library/react";
import { renderWithProviders } from "../../utils/test-utils";
import AddSiteCategoryFormDialog from "@/components/site-category/AddSiteCategoryFormDialog";
import { server } from "../../mocks/server";
import { rest } from "msw";

describe("AddSiteCategoryFormDialog", () => {
  const handleClose = jest.fn();
  it("labelを入力できる", () => {
    renderWithProviders(
      <AddSiteCategoryFormDialog open={true} handleClose={handleClose} />
    );

    const input = screen.getByRole("textbox", {
      name: "Site Category Label",
    }) as HTMLInputElement;

    fireEvent.change(input, { target: { value: "入力したラベル" } });

    expect(input.value).toEqual("入力したラベル");
  });

  it("label入力をしてsubmitすると正しくapiが呼ばれること", async () => {
    const mockFn = jest.fn();
    const BACKEND_ADMIN_API_URL: string =
      process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

    server.use(
      rest.post(
        BACKEND_ADMIN_API_URL + "/site-categories",
        async (req, res, ctx) => {
          const { label } = await req.json();

          mockFn({ label });

          return res(
            ctx.status(200),
            ctx.json({
              id: 2,
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

    renderWithProviders(
      <AddSiteCategoryFormDialog open={true} handleClose={handleClose} />
    );

    const input = screen.getByRole("textbox", {
      name: "Site Category Label",
    }) as HTMLInputElement;

    fireEvent.change(input, { target: { value: "入力したラベル" } });

    const submitButton = screen.getByRole("button", { name: /save/i });
    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() => {
      expect(mockFn).toHaveBeenCalledTimes(1);
      expect(mockFn).toBeCalledWith({ label: "入力したラベル" });
      expect(handleClose).toBeCalled();
    });
  });
});
