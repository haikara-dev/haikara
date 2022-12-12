import { act, fireEvent, screen, waitFor } from "@testing-library/react";
import EditSiteForm from "@/components/site/EditSiteForm";
import { renderWithProviders } from "../../utils/test-utils";
import { server } from "../../mocks/server";
import { rest } from "msw";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const createSite = () => {
  const site = {
    id: 1,
    name: "サイト",
    url: "https://a.b",
    feed_url: "https://a.b/feed",
    active: true,
    cannot_crawl_at: "",
    cannot_crawl: false,
    site_crawl_rule: {
      article_selector: "article",
      title_selector: "h1",
      link_selector: "a",
      description_selector: "p",
      has_data_to_list: true,
      date_selector: "time",
      date_layout: "YYYY-MM-DD",
      is_time_humanize: false,
      is_spa: false,
    },
    site_categories: [
      {
        id: 1,
        label: "カテゴリA",
      },
    ],
  };
  return site;
};

describe("EditSiteForm", () => {
  it("送信のテスト", async () => {
    const mockFn = jest.fn();
    const goBackPageFn = jest.fn();

    server.use(
      rest.put(BACKEND_ADMIN_API_URL + "/sites/:id", async (req, res, ctx) => {
        const { id } = req.params;
        const { name } = await req.json();

        mockFn({
          id,
          name,
        });

        return res(
          ctx.status(200),
          ctx.json({
            id,
            name: name,
            url: "https://nonno.hpplus.jp/fashion",
            active: true,
            created_at: "2022-10-23T16:29:46+09:00",
            updated_at: "2022-12-12T10:54:32+09:00",
            edges: {
              site_crawl_rule: {
                id: 1,
                article_selector: "article",
                title_selector: "h1",
                link_selector: "a",
                has_data_to_list: true,
                date_selector: "time",
                date_layout: "YYYY-MM-DD",
                is_time_humanize: false,
                is_spa: false,
                created_at: "2022-10-23T16:29:46+09:00",
                updated_at: "2022-12-12T10:54:32+09:00",
                edges: {},
              },
            },
          })
        );
      })
    );

    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />, {
      router: {
        back: goBackPageFn,
      },
    });

    const input = screen.getByRole("textbox", { name: "Site Name" });
    fireEvent.change(input, { target: { value: "変更後の名前" } });

    const submitButton = screen.getByRole("button", { name: /更新/i });
    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() => {
      expect(mockFn).toHaveBeenCalledTimes(1);
      expect(mockFn).toBeCalledWith({ id: "1", name: "変更後の名前" });
      expect(goBackPageFn).toHaveBeenCalledTimes(1);
    });
  });
  it.todo("バリデーションのテスト");
});

describe("title_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "Site Name" })).toHaveValue(
      "サイト"
    );
  });
});

describe("title_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "Site URL" })).toHaveValue(
      "https://a.b"
    );
  });
});

describe("title_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "Feed URL" })).toHaveValue(
      "https://a.b/feed"
    );
  });
});

describe("article_selector", () => {
  it("site propsからの初期値が反映する", async () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(await screen.findByLabelText("メディア")).toBeInTheDocument();
  });

  it.todo("Site Categoryのデフォルトチェックのテスト");
});

describe("article_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("textbox", { name: "article_selector" })
    ).toHaveValue("article");
  });
});

describe("title_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "title_selector" })).toHaveValue(
      "h1"
    );
  });
});

describe("link_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "link_selector" })).toHaveValue(
      "a"
    );
  });
});

describe("description_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("textbox", { name: "description_selector" })
    ).toHaveValue("p");
  });
});

describe("has_data_to_list", () => {
  it("site propsからの初期値が反映する true", async () => {
    const site = createSite();
    site.site_crawl_rule.has_data_to_list = true;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /has_data_to_list/ })
    ).toBeChecked();
  });

  it("site propsからの初期値が反映する false", async () => {
    const site = createSite();
    site.site_crawl_rule.has_data_to_list = false;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /has_data_to_list/ })
    ).not.toBeChecked();
  });
});

describe("date_selector", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "date_selector" })).toHaveValue(
      "time"
    );
  });
});

describe("date_layout", () => {
  it("site propsからの初期値が反映する", () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("textbox", { name: "date_layout" })).toHaveValue(
      "YYYY-MM-DD"
    );
  });
});

describe("is_time_humanize", () => {
  it("site propsからの初期値が反映する true", async () => {
    const site = createSite();
    site.site_crawl_rule.is_time_humanize = true;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /is_time_humanize/ })
    ).toBeChecked();
  });

  it("site propsからの初期値が反映する false", async () => {
    const site = createSite();
    site.site_crawl_rule.is_time_humanize = false;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /is_time_humanize/ })
    ).not.toBeChecked();
  });
});

describe("is_spa", () => {
  it("site propsからの初期値が反映する true", async () => {
    const site = createSite();
    site.site_crawl_rule.is_spa = true;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("checkbox", { name: /is_spa/ })).toBeChecked();
  });

  it("site propsからの初期値が反映する false", async () => {
    const site = createSite();
    site.site_crawl_rule.is_spa = false;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(screen.getByRole("checkbox", { name: /is_spa/ })).not.toBeChecked();
  });
});

describe("cannot_crawl", () => {
  it("site propsからの初期値が反映する true", async () => {
    const site = createSite();
    site.cannot_crawl = true;
    site.cannot_crawl_at = new Date().toISOString();
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /cannot_crawl/ })
    ).toBeChecked();
  });

  it("site propsからの初期値が反映する false", async () => {
    const site = createSite();
    site.cannot_crawl = false;
    renderWithProviders(<EditSiteForm site={site} />);
    expect(
      screen.getByRole("checkbox", { name: /cannot_crawl/ })
    ).not.toBeChecked();
  });
});
