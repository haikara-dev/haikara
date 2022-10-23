import { NextPageWithLayout } from "@/pages/_app";
import React, { ReactElement, useEffect } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import Typography from "@mui/material/Typography";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { useRouter } from "next/router";
import {
  useRunGetOGPImageOfArticleMutation,
  useLazyGetArticlesQuery,
  useLazyGetSiteWithSiteCrawlRuleQuery,
  useDeleteArticleMutation,
} from "@/services/adminApi";
import Stack from "@mui/material/Stack";
import LabeledText from "@/components/ui/LabeledText";
import Highlight, { defaultProps } from "prism-react-renderer";
import Link from "next/link";

import Button from "@mui/material/Button";
import ImgproxyImage from "@/components/ImgproxyImage";
import IconButton from "@mui/material/IconButton";
import DeleteIcon from "@mui/icons-material/Delete";

const Site: NextPageWithLayout = () => {
  const router = useRouter();
  const { id } = router.query;
  const [getSiteWithSiteCrawlRule, { data: site }] =
    useLazyGetSiteWithSiteCrawlRuleQuery();

  const [getArticles, { data: articles }] = useLazyGetArticlesQuery();

  const [runGetOGPImageOfArticle] = useRunGetOGPImageOfArticleMutation();

  const [deleteArticle] = useDeleteArticleMutation();

  useEffect(() => {
    if (id) {
      getSiteWithSiteCrawlRule(Number(id));
      getArticles({ site_id: Number(id) });
    }
  }, [id]);

  const onClickDeleteHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    await deleteArticle(id);
  };

  const onClickGetOGPImageHandler = (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    runGetOGPImageOfArticle(id);
  };

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
          <div>active: {site.active ? "ture" : "false"}</div>
          <div>cannot_crawl: {site.cannot_crawl ? "ture" : "false"}</div>
        </div>
      )}
      <h2>Site Crawl Rule</h2>
      {site && site.site_crawl_rule ? (
        <div>
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
          <div>
            has_date_to_list:{" "}
            {site.site_crawl_rule.has_data_to_list ? "ture" : "false"}
          </div>
          <div>
            is_time_humanize:{" "}
            {site.site_crawl_rule.is_time_humanize ? "ture" : "false"}
          </div>
          <Stack direction="row" gap={2}>
            <LabeledText
              label="date_selector"
              value={site.site_crawl_rule.date_selector}
            />
            <LabeledText
              label="date_layout"
              value={site.site_crawl_rule.date_layout}
            />
          </Stack>

          <div>is_spa: {site.site_crawl_rule.is_spa ? "ture" : "false"}</div>
        </div>
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
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>OGP</TableCell>
                  <TableCell>OGP Image</TableCell>
                  <TableCell>Published</TableCell>
                  <TableCell>Article Title</TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {articles.data.map((article) => {
                  return (
                    <TableRow key={article.id}>
                      <TableCell>
                        <Button
                          onClick={onClickGetOGPImageHandler.bind(
                            this,
                            article.id
                          )}
                        >
                          OGP
                        </Button>
                      </TableCell>
                      <TableCell>
                        {article.ogp_image_url && (
                          <div
                            style={{
                              display: "inline-block",
                              width: "200px",
                              height: "105px",
                              position: "relative",
                            }}
                          >
                            <ImgproxyImage
                              src={article.ogp_image_url}
                              width={200 * 2}
                              height={105 * 2}
                              objectFit="contain"
                            />
                          </div>
                        )}
                      </TableCell>
                      <TableCell>
                        {new Date(article.published_at).toLocaleDateString()}
                      </TableCell>
                      <TableCell>{article.title}</TableCell>
                      <TableCell>
                        <Button
                          component="a"
                          href={article.url}
                          target="_blank"
                          rel="noopener noreferrer"
                        >
                          Read
                        </Button>
                      </TableCell>
                      <TableCell>
                        <IconButton
                          onClick={onClickDeleteHandler.bind(this, article.id)}
                          aria-label="remove"
                        >
                          <DeleteIcon />
                        </IconButton>
                      </TableCell>
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          </TableContainer>
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
