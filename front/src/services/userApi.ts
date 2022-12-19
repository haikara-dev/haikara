import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

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

export type ListResponse<T> = {
  totalCount: number;
  totalPage: number;
  pageSize: number;
  data: T[];
};

export type GetArticlesArg = {
  page?: number;
  site_id?: number;
};

// Define a service using a base URL and expected endpoints
export const userApi = createApi({
  reducerPath: "userApi",
  baseQuery: fetchBaseQuery({
    baseUrl: BACKEND_API_URL,
    prepareHeaders: async (headers, { getState }) => {
      headers.set("Content-Type", "application/json");
      return headers;
    },
  }),
  tagTypes: ["Articles"],
  endpoints: (builder) => ({
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
  }),
});

/*
  Hooks
 */
export const { useLazyGetArticlesQuery } = userApi;
