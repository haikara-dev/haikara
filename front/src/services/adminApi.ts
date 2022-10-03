import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

import { RootState } from "@/app/store";
import { User } from "@/features/auth/authSlice";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export type Article = {
  id: number;
  title: string;
  url: string;
  published_at: string;
};

export type Feed = {
  id: number;
  count: number;
  created_at: string;
  indexed_at: string;
  site_id: number;
  site_name: string;
};

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
  site_crawl_rule?: SiteCrawlRule;
};

export type NestedSiteWithSiteCrawlRuleServerResponse = Site & {
  edges?: {
    site_crawl_rule: SiteCrawlRule;
  };
};

export type DryRunResult = {
  count: number;
  contents: string;
};

export type ListResponse<T> = {
  totalCount: number;
  totalPage: number;
  pageSize: number;
  data: T[];
};

export type DeleteResponse = {
  massage: string;
};

export type UpdateUserRoleArg = {
  id: number;
  body: {
    role: string;
  };
};

export type AddSiteArg = {
  body: {
    site: SiteWithSiteCrawlRule;
  };
};

export type UpdateSiteArg = {
  id: number;
  body: {
    site: SiteWithSiteCrawlRule;
  };
};

export type ActiveSiteArg = {
  id: number;
  body: {
    active: boolean;
  };
};

export type DeActiveSiteArg = {
  id: number;
  body: {
    active: boolean;
  };
};

export type RunSiteCrawlingResponse = {
  id: number;
  url: string;
  rss: string;
};

export type DryRunSiteCrawlingResponse = {
  count: number;
  contents: string;
};

export type GetSiteRssUrlResponse = {
  url: string;
};

export type GetSiteRssUrlByUrlResponse = {
  url: string;
};

export const adminApi = createApi({
  reducerPath: "adminApi",
  baseQuery: fetchBaseQuery({
    baseUrl: BACKEND_ADMIN_API_URL,
    prepareHeaders: async (headers, { getState }) => {
      const authUser = (getState() as RootState).auth.authUser;

      if (authUser) {
        const idToken = await authUser.getIdToken();
        if (idToken) {
          headers.set("Authorization", `Bearer ${idToken}`);
        }
      }
      headers.set("Content-Type", "application/json");
      return headers;
    },
  }),
  tagTypes: ["Sites"],
  endpoints: (builder) => ({
    /*
        Article
     */
    getArticles: builder.query<ListResponse<Article>, number | void>({
      query: (page = 1) => ({
        url: `/articles?page=${page}`,
      }),
    }),
    deleteArticle: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/articles/${id}`,
        method: "DELETE",
      }),
    }),
    /*
        User
     */
    getUsers: builder.query<ListResponse<User>, number | void>({
      query: (page = 1) => ({
        url: `/users?page=${page}`,
      }),
    }),
    updateUserRole: builder.mutation<User, UpdateUserRoleArg>({
      query: (queryArg) => ({
        url: `/users/role/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
    }),
    /*
       Feed
    */
    getFeeds: builder.query<ListResponse<Feed>, number | void>({
      query: (page = 1) => ({
        url: `/feeds/lite?page=${page}`,
      }),
    }),
    deleteFeed: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/feeds/${id}`,
        method: "DELETE",
      }),
    }),
    runParseFeed: builder.mutation<Feed, number>({
      query: (id) => ({
        url: `/feeds/parse/${id}`,
      }),
    }),
    /*
       Site
    */
    getSites: builder.query<ListResponse<Site>, number | void>({
      // TODO: site.cannot_crawl = site.cannot_crawl_at ? true : false;
      query: (page = 1) => ({
        url: `/sites?page=${page}`,
      }),
      transformResponse: (response: ListResponse<Site>) =>
        addCanCrawlFieldToSiteListResponse(response),
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({ type: "Sites" as const, id })),
              { type: "Sites", id: "LIST" },
            ]
          : [{ type: "Sites", id: "LIST" }],
    }),
    getSiteWithSiteCrawlRule: builder.query<SiteWithSiteCrawlRule, number>({
      // TODO: site.cannot_crawl = site.cannot_crawl_at ? true : false;
      query: (id) => ({
        url: `/sites/${id}`,
      }),
      transformResponse: (
        response: NestedSiteWithSiteCrawlRuleServerResponse
      ) => {
        return addCanCrawlFieldToSite(
          unNestNestedSiteWithSiteCrawlRuleServerResponse(response)
        ) as SiteWithSiteCrawlRule;
      },
    }),
    addSite: builder.mutation<Site, AddSiteArg>({
      query: (queryArg) => ({
        url: `/sites`,
        method: "POST",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
    }),
    updateSite: builder.mutation<Site, UpdateSiteArg>({
      query: (queryArg) => ({
        url: `/sites/${queryArg.id}`,
        method: "PUT",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
    }),
    deleteSite: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/sites/${id}`,
        method: "DELETE",
      }),
    }),
    activeSite: builder.mutation<Site, ActiveSiteArg>({
      query: (queryArg) => ({
        url: `/sites/active/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, arg) => [{ type: "Sites", id: arg.id }],
    }),
    deActiveSite: builder.mutation<Site, DeActiveSiteArg>({
      query: (queryArg) => ({
        url: `/sites/deActive/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, arg) => [{ type: "Sites", id: arg.id }],
    }),
    runSiteCrawling: builder.mutation<RunSiteCrawlingResponse, number>({
      query: (id) => ({
        url: `/sites/run-crawling/${id}`,
        method: "GET",
      }),
    }),
    dryRunSiteCrawling: builder.mutation<DryRunSiteCrawlingResponse, number>({
      query: (id) => ({
        url: `/sites/dry-run-crawling/${id}`,
        method: "GET",
      }),
    }),
    getSiteRssUrl: builder.mutation<GetSiteRssUrlResponse, number>({
      query: (id) => ({
        url: `/sites/get-rss-url/${id}`,
        method: "GET",
      }),
    }),
    getSiteRssUrlByUrl: builder.mutation<GetSiteRssUrlByUrlResponse, string>({
      query: (url) => ({
        url: `/sites/get-rss-url-by-url`,
        params: { url: url },
        method: "GET",
      }),
    }),
  }),
});

/*
  transformResponse
 */

/**
 * Add can_crawl field to SiteListResponse
 * @param response
 */
const addCanCrawlFieldToSiteListResponse = (response: ListResponse<Site>) => {
  response.data.forEach((site) => {
    site.cannot_crawl = site.cannot_crawl_at ? true : false;
  });
  return response;
};

/**
 * Add can_crawl field to Site
 * @param site
 */
const addCanCrawlFieldToSite = (site: Site) => {
  site.cannot_crawl = site.cannot_crawl_at ? true : false;
  return site;
};

/**
 * UnNest NestedSiteWithSiteCrawlRuleServerResponse
 *  {edges.site_crawl_rule} -> {site_crawl_rule}
 * @param response
 */
const unNestNestedSiteWithSiteCrawlRuleServerResponse = (
  response: NestedSiteWithSiteCrawlRuleServerResponse
) => {
  const site_crawl_rule = response.edges?.site_crawl_rule;
  delete response.edges;
  const parsedResponse = {
    ...response,
  } as SiteWithSiteCrawlRule;
  parsedResponse.site_crawl_rule = site_crawl_rule;
  return parsedResponse;
};

/*
  Hooks
 */
export const { useGetArticlesQuery, useDeleteArticleMutation } = adminApi;

export const { useGetUsersQuery, useUpdateUserRoleMutation } = adminApi;

export const {
  useGetFeedsQuery,
  useDeleteFeedMutation,
  useRunParseFeedMutation,
} = adminApi;

export const {
  useGetSitesQuery,
  useGetSiteWithSiteCrawlRuleQuery,
  useLazyGetSiteWithSiteCrawlRuleQuery,
  useAddSiteMutation,
  useUpdateSiteMutation,
  useDeleteSiteMutation,
  useActiveSiteMutation,
  useDeActiveSiteMutation,
  useRunSiteCrawlingMutation,
  useDryRunSiteCrawlingMutation,
  useGetSiteRssUrlMutation,
  useGetSiteRssUrlByUrlMutation,
} = adminApi;
