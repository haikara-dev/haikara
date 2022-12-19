import { rest } from "msw";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export const handlers = [
  rest.get(BACKEND_API_URL + "/articles", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        totalCount: 7175,
        totalPage: 72,
        pageSize: 100,
        data: [
          {
            id: 1,
            title: "記事タイトル1",
            url: "https://a.b/articles/1",
            published_at: "2022-11-15T20:46:54+09:00",
            ogp_image_url:
              "http://nginx_container/uploads/ogp_images/2022/11/15/1.jpg",
            site: {
              id: 50,
              name: "fashion news",
              url: "https://a.b/",
            },
          },
          {
            id: 2,
            title: "記事タイトル2",
            url: "https://a.b/articles/2",
            published_at: "2022-11-15T20:46:54+09:00",
            ogp_image_url:
              "http://nginx_container/uploads/ogp_images/2022/11/15/2.jpg",
            site: {
              id: 50,
              name: "fashion news",
              url: "https://a.b/",
            },
          },
        ],
      })
    );
  }),
];
