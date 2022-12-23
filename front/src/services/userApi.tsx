import useSWR from "swr";

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

export const useGetArticlesQuery = (queryArg: GetArticlesArg) => {
  const fetcher = (url: string, queryArg: GetArticlesArg): Promise<any> => {
    const page = queryArg.page ? queryArg.page : 1;
    const query = new URLSearchParams({ page: page.toString() }).toString();

    const headers = new Headers();
    headers.set("Content-Type", "application/json");

    return fetch(`${BACKEND_API_URL}${url}?${query}`, {
      method: "GET",
      headers: headers,
    }).then((res) => res.json());
  };

  const { data, error, isLoading } = useSWR<ListResponse<Article>>(
    ["/articles", queryArg],
    ([url, queryArg]) => fetcher(url, queryArg)
  );
  return { data, error, isLoading };
};
