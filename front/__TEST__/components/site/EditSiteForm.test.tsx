import { screen } from "@testing-library/react";
import EditSiteForm from "@/components/site/EditSiteForm";
import { renderWithProviders } from "../../utils/test-utils";

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
  it("site propsからの初期値が反映する", async () => {
    const site = createSite();
    renderWithProviders(<EditSiteForm site={site} />);
    // Site
    expect(screen.getByRole("textbox", { name: "Site Name" })).toHaveValue(
      "サイト"
    );
    expect(screen.getByRole("textbox", { name: "Site URL" })).toHaveValue(
      "https://a.b"
    );
    expect(screen.getByRole("textbox", { name: "Feed URL" })).toHaveValue(
      "https://a.b/feed"
    );
    // Site Category
    expect(await screen.findByLabelText("メディア")).toBeInTheDocument();
    // Site Crawl Rule
    expect(
      screen.getByRole("textbox", { name: "article_selector" })
    ).toHaveValue("article");
    expect(screen.getByRole("textbox", { name: "title_selector" })).toHaveValue(
      "h1"
    );
    expect(screen.getByRole("textbox", { name: "link_selector" })).toHaveValue(
      "a"
    );
    expect(
      screen.getByRole("textbox", { name: "description_selector" })
    ).toHaveValue("p");
    // todo: has_data_to_list
    expect(
      screen.getByRole("checkbox", { name: "has_data_to_list" })
    ).toBeChecked();
    expect(screen.getByRole("textbox", { name: "date_selector" })).toHaveValue(
      "time"
    );
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

describe("EditSiteForm", () => {
  it.todo("Site Categoryのデフォルトチェックのテスト");
  it.todo("送信のテスト");
  it.todo("バリデーションのテスト");
});
