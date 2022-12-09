import { act, fireEvent, screen, waitFor } from "@testing-library/react";

import SiteRow from "@/components/site/SiteRow";
import { renderWithProviders } from "../../utils/test-utils";
import TableBody from "@mui/material/TableBody";
import Table from "@mui/material/Table";

describe("SiteRow", () => {
  it("renders", () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "https://a.b/feed",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: false,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    expect(
      screen.getByRole("button", {
        name: /run/i,
      })
    ).toBeInTheDocument();

    expect(
      screen.getByRole("button", {
        name: /dry/i,
      })
    ).toBeInTheDocument();

    expect(screen.getByText("サイト")).toBeInTheDocument();
    expect(screen.getByText("https://a.b")).toBeInTheDocument();
    expect(screen.getByText("カテゴリA")).toBeInTheDocument();
  });

  it("feed_urlがあれば、「RSS」と表示する", () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "https://a.b/feed",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: false,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    expect(screen.getByTestId("have-feed-url")).toHaveTextContent("RSS");
  });
  it("feed_urlがなければ、「-」と表示する", () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: false,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    expect(screen.getByTestId("have-feed-url")).toHaveTextContent("-");
  });

  it("cannot_crawlがtrueであれば背景を赤にする", () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: true,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    expect(screen.getByRole("row")).toHaveStyle({ backgroundColor: "#f5c4c4" });
  });

  it("cannot_crawlがtrueであれば背景を白にする", () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: false,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    expect(screen.getByRole("row")).toHaveStyle({ backgroundColor: "white" });
  });

  it("DryボタンクリックでopenDryDialogが呼ばれる", async () => {
    const site = {
      id: 1,
      name: "サイト",
      url: "https://a.b",
      feed_url: "",
      active: true,
      cannot_crawl_at: "",
      cannot_crawl: false,
      site_categories: [
        {
          id: 1,
          label: "カテゴリA",
        },
      ],
    };

    const openDryDialog = jest.fn();

    renderWithProviders(
      <Table>
        <TableBody>
          <SiteRow site={site} openDryDialog={openDryDialog} />
        </TableBody>
      </Table>
    );

    await act(() => {
      fireEvent.click(screen.getByRole("button", { name: /dry/i }));
    });
    await waitFor(() => {
      expect(openDryDialog).toHaveBeenCalledTimes(1);
    });
  });
});
