import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

import { RootState } from "@/app/store";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export type Article = {
  id: number;
  title: string;
  url: string;
  published_at: string;
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
  endpoints: (builder) => ({
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
  }),
});

/*
  Hooks
 */
export const { useGetArticlesQuery, useDeleteArticleMutation } = adminApi;
