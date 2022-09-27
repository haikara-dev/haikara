export type Site = {
  id: number;
  name: string;
  url: string;
  feed_url: string;
  active: boolean;
  cannot_crawl_at: string;
  cannot_crawl: boolean;
};

export type SiteCrawlRule = {
  article_selector: string;
  title_selector: string;
  link_selector: string;
  description_selector: string;
  has_data_to_list: boolean;
  date_selector: string;
  date_layout: string;
  is_time_humanize: boolean;
  is_spa: boolean;
};

export type SiteWithSiteCrawlRule = Site & {
  site_crawl_rule: SiteCrawlRule;
};

export type DryRunResult = {
  count: number;
  contents: string;
};
