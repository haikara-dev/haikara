import { NextPageWithLayout } from "@/pages/_app";
import React, { ReactElement, useEffect } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import Typography from "@mui/material/Typography";
import { useRouter } from "next/router";
import {
  useLazyGetArticlesQuery,
  useLazyGetSiteWithSiteCrawlRuleQuery,
} from "@/services/adminApi";
import Stack from "@mui/material/Stack";
import LabeledText from "@/components/ui/LabeledText";
import Highlight, { defaultProps } from "prism-react-renderer";
import Link from "next/link";

const Site: NextPageWithLayout = () => {
  const router = useRouter();
  const { id } = router.query;
  const [getSiteWithSiteCrawlRule, { data: site }] =
    useLazyGetSiteWithSiteCrawlRuleQuery();

  const [getArticles, { data: articles }] = useLazyGetArticlesQuery();
  useEffect(() => {
    if (id) {
      getSiteWithSiteCrawlRule(Number(id));
      getArticles({ site_id: Number(id) });
    }
  }, [id]);
  return (
    <div>
      <Typography variant="h3" component="h1">
        {site && site.name}
      </Typography>
      <h2>Site</h2>
      {site && (
        <div>
          <div>
            url:{" "}
            <a href={site.url} target="_blank" rel="noopener noreferrer">
              {site.url}
            </a>
          </div>
          <div>feed_url: {site.feed_url}</div>
          <div>active: {site.active}</div>
          <div>cannot_crawl: {site.cannot_crawl}</div>
        </div>
      )}
      <h2>Site Crawl Rule</h2>
      {site && site.site_crawl_rule ? (
        <Stack direction="row" gap={2}>
          <LabeledText
            label="article_selector"
            value={site.site_crawl_rule.article_selector}
          />
          <LabeledText
            label="title_selector"
            value={site.site_crawl_rule.title_selector}
          />
          <LabeledText
            label="link_selector"
            value={site.site_crawl_rule.link_selector}
          />
        </Stack>
      ) : (
        <div>none</div>
      )}
      {site && (
        <Highlight
          {...defaultProps}
          code={JSON.stringify(site, null, 2)}
          language="json"
        >
          {({ className, style, tokens, getLineProps, getTokenProps }) => (
            <pre className={className} style={style}>
              {tokens.map((line, i) => (
                <div key={i} {...getLineProps({ line })}>
                  {line.map((token, key) => (
                    <span key={key} {...getTokenProps({ token })} />
                  ))}
                </div>
              ))}
            </pre>
          )}
        </Highlight>
      )}
      <h2>Articles</h2>
      {articles && articles.data ? (
        <>
          <div>{articles.totalCount}件</div>
          <ul>
            {articles.data.map((article) => (
              <li key={article.id}>
                <a href={article.url} target="_blank" rel="noopener noreferrer">
                  {article.title}
                </a>
              </li>
            ))}
          </ul>
          {articles.totalCount > articles.data.length && (
            <div>
              <Link href={`/article?site_id=${id}`}>more...</Link>
            </div>
          )}
        </>
      ) : (
        <div>0件</div>
      )}
    </div>
  );
};
Site.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default Site;
