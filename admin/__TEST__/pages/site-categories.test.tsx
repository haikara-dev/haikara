import {
  act,
  fireEvent,
  screen,
  waitFor,
  waitForElementToBeRemoved,
} from "@testing-library/react";
import "@testing-library/jest-dom";
import { mockAuthorizedAuth, renderWithProviders } from "../utils/test-utils";
import SiteCategories from "@/pages/site-categories";
import { rest } from "msw";
import { server } from "../mocks/server";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("SiteCategories", () => {
  it("render heading", () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
    });
    const heading = screen.getByText(/Site Category/i);
    expect(heading).toBeInTheDocument();
  });

  it("render heading with Login(role user)", async () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });

    const heading = screen.getByText(/Site Category/i);
    expect(heading).toBeInTheDocument();

    const loading = screen.getByText(/Loading/i);
    expect(loading).toBeInTheDocument();

    const articles = await screen.findByText(/メディア/i);
    expect(articles).toBeInTheDocument();
  });

  it("削除ボタンをクリックすると確認ダイアログが表示させる", async () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });

    const removeButton = await screen.findByRole("button", {
      name: "remove",
    });

    expect(removeButton).toBeInTheDocument();

    await act(() => {
      removeButton.click();
    });

    let confirmDialog;
    await waitFor(() => {
      confirmDialog = screen.getByRole("dialog", {
        name: /カテゴリの削除/i,
      });
    });

    expect(confirmDialog).toBeInTheDocument();
    expect(confirmDialog).toBeVisible();
  });

  it("削除確認ダイアログでキャンセルボタンをクリックするとダイヤログが閉じる", async () => {
    renderWithProviders(<SiteCategories />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });

    const removeButton = await screen.findByRole("button", {
      name: "remove",
    });

    expect(removeButton).toBeInTheDocument();

    await act(() => {
      removeButton.click();
    });

    let confirmDialog;
    await waitFor(() => {
      confirmDialog = screen.getByRole("dialog", {
        name: /カテゴリの削除/i,
      });
    });

    await act(() => {
      fireEvent.click(screen.getByRole("button", { name: "キャンセル" }));
    });

    expect(confirmDialog).not.toBeVisible();
  });

  it("削除確認ダイアログで削除ボタンをクリックすると削除apiが呼ばれる", async () => {
    const apiHookFn = jest.fn();

    server.use(
      rest.delete(
        BACKEND_ADMIN_API_URL + "/site-categories/:id",
        async (req, res, ctx) => {
          apiHookFn();
          return res(ctx.status(200), ctx.json({ message: "deleted" }));
        }
      )
    );

    renderWithProviders(<SiteCategories />, {
      router: {},
      preloadedState: { ...mockAuthorizedAuth("admin") },
    });

    const removeButton = await screen.findByRole("button", {
      name: "remove",
    });

    expect(removeButton).toBeInTheDocument();

    await act(() => {
      removeButton.click();
    });

    let confirmDialog;
    await waitFor(() => {
      confirmDialog = screen.getByRole("dialog", {
        name: /カテゴリの削除/i,
      });
    });

    await act(() => {
      fireEvent.click(screen.getByRole("button", { name: "削除" }));
    });

    await waitFor(() => {
      expect(apiHookFn).toHaveBeenCalledTimes(1);
    });

    await waitForElementToBeRemoved(() =>
      screen.getByRole("dialog", {
        name: /カテゴリの削除/i,
      })
    );

    expect(
      screen.queryByRole("dialog", {
        name: /カテゴリの削除/i,
      })
    ).not.toBeInTheDocument();
  });
});
