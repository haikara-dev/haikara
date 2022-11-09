import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

import { RootState } from "@/app/store";
import { User } from "@/features/auth/authSlice";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export type ArticleSite = {
  id: number;
  name: string;
  url: string;
};

export type Article = {
  id: number;
  title: string;
  url: string;
  published_at: string;
  ogp_image_url: string;
  site: ArticleSite;
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

export type SiteWithCategory = Site & {
  site_categories: SiteCategory[];
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

export type SiteWithCrawlRuleAndCategory = Site & {
  site_crawl_rule: SiteCrawlRule;
  site_categories: SiteCategory[];
};

export type NestedSiteWithCrawlRuleAndCategoryServerResponse = Site & {
  edges?: {
    site_crawl_rule: SiteCrawlRule;
    site_categories: SiteCategory[];
  };
};

export type SiteCategory = {
  id: number;
  label: string;
  sites_count?: number;
};

export type GetSiteCategoriesArg = {
  page?: number;
};

export type AddSiteCategoryArg = {
  body: {
    label: string;
  };
};

export type UpdateSiteCategoryArg = {
  id: number;
  body: {
    label: string;
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

export type GetArticlesArg = {
  page?: number;
  site_id?: number;
};

export type RunGetOGPImageOfArticleResponse = {
  massage: string;
};

export type UpdateUserRoleArg = {
  id: number;
  body: {
    role: string;
  };
};

export type AddSiteArg = {
  body: Site & {
    site_crawl_rule: SiteCrawlRule;
    site_category_ids: number[];
  };
};

export type UpdateSiteArg = {
  id: number;
  body: Site & {
    site_crawl_rule: SiteCrawlRule;
    site_category_ids: number[];
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

export type SiteImportExport = {
  name: string;
  url: string;
  feed_url: string;
  active: boolean;
  cannot_crawl_at: string;
};

type SiteCrawlRuleImportExport = {
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

export type SiteWithSiteCrawlRuleImportExport = SiteImportExport & {
  site_crawl_rule?: SiteCrawlRuleImportExport;
};

export type ImportSiteArg = {
  form: FormData;
};

export type ImportSiteResponse = {
  count: string;
  reqSites: SiteWithCrawlRuleAndCategory[];
};

export type AdminDashboard = {
  siteSize: number;
  articleSize: number;
  feedSize: number;
  userSize: number;
  ogpImageSize: number;
};
export const adminApi = createApi({
  reducerPath: "adminApi",
  baseQuery: fetchBaseQuery({
    baseUrl: BACKEND_ADMIN_API_URL,
    prepareHeaders: async (headers, { getState, endpoint }) => {
      const authUser = (getState() as RootState).auth.authUser;

      if (authUser) {
        const idToken = await authUser.getIdToken();
        if (idToken) {
          headers.set("Authorization", `Bearer ${idToken}`);
        }
      }

      headers.set("Content-Type", "application/json");

      if (endpoint === "importSite") {
        headers.delete("Content-Type");
      }

      return headers;
    },
  }),
  tagTypes: [
    "Sites",
    "SiteCategories",
    "Feeds",
    "Articles",
    "Users",
    "SiteWithCrawlRuleAndCategory",
    "AdminDashboard",
  ],
  endpoints: (builder) => ({
    /*
        Article
     */
    getArticles: builder.query<ListResponse<Article>, GetArticlesArg>({
      query: (queryArg) => {
        const page = queryArg.page ? queryArg.page : 1;

        return {
          url: `/articles`,
          params: { ...queryArg, page },
        };
      },
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({
                type: "Articles" as const,
                id,
              })),
              { type: "Articles", id: "PARTIAL-LIST" },
            ]
          : [{ type: "Articles", id: "PARTIAL-LIST" }],
    }),
    deleteArticle: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/articles/${id}`,
        method: "DELETE",
      }),
      invalidatesTags: (result, error, id) => [
        { type: "Articles", id },
        { type: "Articles", id: "PARTIAL-LIST" },
      ],
    }),
    runGetOGPImageOfArticle: builder.mutation<
      RunGetOGPImageOfArticleResponse,
      number
    >({
      query: (id) => ({
        url: `/articles/run-get-ogp-image/${id}`,
        method: "GET",
      }),
      invalidatesTags: (result, error, id) => [
        { type: "Articles", id },
        { type: "Articles", id: "PARTIAL-LIST" },
      ],
    }),
    /*
        User
     */
    getUsers: builder.query<ListResponse<User>, number | void>({
      query: (page = 1) => ({
        url: `/users?page=${page}`,
      }),
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({
                type: "Users" as const,
                id,
              })),
              { type: "Users", id: "PARTIAL-LIST" },
            ]
          : [{ type: "Users", id: "PARTIAL-LIST" }],
    }),
    updateUserRole: builder.mutation<User, UpdateUserRoleArg>({
      query: (queryArg) => ({
        url: `/users/role/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
      invalidatesTags: (result, error, queryArg) => [
        { type: "Users", id: queryArg.id },
        { type: "Users", id: "PARTIAL-LIST" },
      ],
    }),
    /*
       Feed
    */
    getFeeds: builder.query<ListResponse<Feed>, number | void>({
      query: (page = 1) => ({
        url: `/feeds/lite?page=${page}`,
      }),
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({ type: "Feeds" as const, id })),
              { type: "Feeds", id: "PARTIAL-LIST" },
            ]
          : [{ type: "Feeds", id: "PARTIAL-LIST" }],
    }),
    deleteFeed: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/feeds/${id}`,
        method: "DELETE",
      }),
      invalidatesTags: (result, error, id) => [
        { type: "Feeds", id },
        { type: "Feeds", id: "PARTIAL-LIST" },
      ],
    }),
    runParseFeed: builder.mutation<Feed, number>({
      query: (id) => ({
        url: `/feeds/parse/${id}`,
      }),
      invalidatesTags: (result, error, id) => [{ type: "Articles" }],
    }),
    /*
       Site
    */
    getSites: builder.query<ListResponse<SiteWithCategory>, number | void>({
      // TODO: site.cannot_crawl = site.cannot_crawl_at ? true : false;
      query: (page = 1) => ({
        url: `/sites?page=${page}`,
      }),
      transformResponse: (response: ListResponse<SiteWithCategory>) =>
        addCanCrawlFieldToSiteListResponse(response),
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({ type: "Sites" as const, id })),
              { type: "Sites", id: "PARTIAL-LIST" },
            ]
          : [{ type: "Sites", id: "PARTIAL-LIST" }],
    }),
    getSiteWithCrawlRuleAndCategory: builder.query<
      SiteWithCrawlRuleAndCategory,
      number
    >({
      // TODO: site.cannot_crawl = site.cannot_crawl_at ? true : false;
      query: (id) => ({
        url: `/sites/${id}`,
      }),
      transformResponse: (
        response: NestedSiteWithCrawlRuleAndCategoryServerResponse
      ) => {
        return addCanCrawlFieldToSite(
          unNestNestedSiteWithCrawlRuleAndCategoryServerResponse(response)
        ) as SiteWithCrawlRuleAndCategory;
      },
      providesTags: (result, error, id) => [
        { type: "SiteWithCrawlRuleAndCategory", id },
      ],
    }),
    addSite: builder.mutation<Site, AddSiteArg>({
      query: (queryArg) => ({
        url: `/sites`,
        method: "POST",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, queryArg) => [{ type: "Sites" }],
    }),
    updateSite: builder.mutation<Site, UpdateSiteArg>({
      query: (queryArg) => ({
        url: `/sites/${queryArg.id}`,
        method: "PUT",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, queryArg) => [
        { type: "Sites", id: queryArg.id },
        { type: "Sites", id: "PARTIAL-LIST" },
        { type: "SiteWithCrawlRuleAndCategory", id: queryArg.id },
      ],
    }),
    deleteSite: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/sites/${id}`,
        method: "DELETE",
      }),
      invalidatesTags: (result, error, id) => [
        { type: "Sites", id },
        { type: "Sites", id: "PARTIAL-LIST" },
        { type: "SiteWithCrawlRuleAndCategory", id },
      ],
    }),
    activeSite: builder.mutation<Site, ActiveSiteArg>({
      query: (queryArg) => ({
        url: `/sites/active/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, arg) => [
        { type: "Sites", id: arg.id },
        { type: "Sites", id: "PARTIAL-LIST" },
        { type: "SiteWithCrawlRuleAndCategory", id: arg.id },
      ],
    }),
    deActiveSite: builder.mutation<Site, DeActiveSiteArg>({
      query: (queryArg) => ({
        url: `/sites/deActive/${queryArg.id}`,
        method: "PATCH",
        body: queryArg.body,
      }),
      transformResponse: (response: Site) => addCanCrawlFieldToSite(response),
      invalidatesTags: (result, error, arg) => [
        { type: "Sites", id: arg.id },
        { type: "Sites", id: "PARTIAL-LIST" },
        { type: "SiteWithCrawlRuleAndCategory", id: arg.id },
      ],
    }),
    runSiteCrawling: builder.mutation<RunSiteCrawlingResponse, number>({
      query: (id) => ({
        url: `/sites/run-crawling/${id}`,
        method: "GET",
      }),
      invalidatesTags: (result, error, arg) => [{ type: "Feeds" }],
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
    exportSite: builder.mutation<SiteWithSiteCrawlRuleImportExport[], void>({
      query: (queryArg) => ({
        url: `/sites/export`,
        method: "GET",
      }),
    }),
    importSite: builder.mutation<ImportSiteResponse, ImportSiteArg>({
      query: (queryArg) => ({
        url: `/sites/import`,
        method: "POST",
        body: queryArg.form,
      }),
    }),
    /*
      Site Category
    */
    getSiteCategories: builder.query<
      ListResponse<SiteCategory>,
      GetSiteCategoriesArg
    >({
      query: (queryArg) => {
        const page = queryArg.page ? queryArg.page : 1;
        return {
          url: `/site-categories`,
          params: { ...queryArg, page },
        };
      },
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({
                type: "SiteCategories" as const,
                id,
              })),
              { type: "SiteCategories", id: "PARTIAL-LIST" },
            ]
          : [{ type: "SiteCategories", id: "PARTIAL-LIST" }],
    }),
    addSiteCategory: builder.mutation<SiteCategory, AddSiteCategoryArg>({
      query: (queryArg) => ({
        url: `/site-categories`,
        method: "POST",
        body: queryArg.body,
      }),
      invalidatesTags: (result, error, queryArg) => [
        { type: "SiteCategories" },
      ],
    }),
    updateSiteCategory: builder.mutation<SiteCategory, UpdateSiteCategoryArg>({
      query: (queryArg) => ({
        url: `/site-categories/${queryArg.id}`,
        method: "PUT",
        body: queryArg.body,
      }),
      invalidatesTags: (result, error, queryArg) => [
        { type: "SiteCategories", id: queryArg.id },
        { type: "SiteCategories", id: "PARTIAL-LIST" },
      ],
    }),
    deleteSiteCategory: builder.mutation<DeleteResponse, number>({
      query: (id) => ({
        url: `/site-categories/${id}`,
        method: "DELETE",
      }),
      invalidatesTags: (result, error, id) => [
        { type: "SiteCategories", id },
        { type: "SiteCategories", id: "PARTIAL-LIST" },
      ],
    }),
    /*
       AdminDashboard
     */
    getAdminDashboard: builder.query<AdminDashboard, void>({
      query: () => ({
        url: `/dashboard`,
      }),
      providesTags: (result) => [{ type: "AdminDashboard" }],
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
const addCanCrawlFieldToSiteListResponse = (
  response: ListResponse<SiteWithCategory>
) => {
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
 * UnNest NestedSiteWithCrawlRuleAndCategoryServerResponse
 *  {edges.site_crawl_rule} -> {site_crawl_rule}
 * @param response
 */
const unNestNestedSiteWithCrawlRuleAndCategoryServerResponse = (
  response: NestedSiteWithCrawlRuleAndCategoryServerResponse
) => {
  const site_crawl_rule = response.edges?.site_crawl_rule;
  const site_categories = response.edges?.site_categories;
  delete response.edges;
  const parsedResponse = {
    ...response,
  } as SiteWithCrawlRuleAndCategory;
  if (site_crawl_rule) {
    parsedResponse.site_crawl_rule = site_crawl_rule;
  } else {
    parsedResponse.site_crawl_rule = {
      article_selector: "",
      title_selector: "",
      link_selector: "",
      description_selector: "",
      has_data_to_list: true,
      date_selector: "",
      date_layout: "",
      is_time_humanize: false,
      is_spa: false,
    };
  }

  if (site_categories) {
    parsedResponse.site_categories = site_categories;
  } else {
    parsedResponse.site_categories = [];
  }

  return parsedResponse;
};

/*
  Hooks
 */
export const {
  useGetArticlesQuery,
  useLazyGetArticlesQuery,
  useDeleteArticleMutation,
  useRunGetOGPImageOfArticleMutation,
} = adminApi;

export const { useGetUsersQuery, useUpdateUserRoleMutation } = adminApi;

export const {
  useGetFeedsQuery,
  useDeleteFeedMutation,
  useRunParseFeedMutation,
} = adminApi;

export const {
  useGetSitesQuery,
  useGetSiteWithCrawlRuleAndCategoryQuery,
  useLazyGetSiteWithCrawlRuleAndCategoryQuery,
  useAddSiteMutation,
  useUpdateSiteMutation,
  useDeleteSiteMutation,
  useActiveSiteMutation,
  useDeActiveSiteMutation,
  useRunSiteCrawlingMutation,
  useDryRunSiteCrawlingMutation,
  useGetSiteRssUrlMutation,
  useGetSiteRssUrlByUrlMutation,
  useExportSiteMutation,
  useImportSiteMutation,
} = adminApi;

export const {
  useGetSiteCategoriesQuery,
  useAddSiteCategoryMutation,
  useUpdateSiteCategoryMutation,
  useDeleteSiteCategoryMutation,
} = adminApi;

export const { useGetAdminDashboardQuery } = adminApi;
