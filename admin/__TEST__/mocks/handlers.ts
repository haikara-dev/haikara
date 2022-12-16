import { rest } from "msw";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export const handlers = [
  rest.get(BACKEND_API_URL + "/users/current", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        id: 7,
        created_at: "2022-09-27T15:20:16+09:00",
        updated_at: "2022-09-27T15:22:59+09:00",
        UUID: "0DrCl7hccfdsdsfwtwtwe",
        email: "a@b",
        role: "user",
      })
    );
  }),

  rest.get(BACKEND_API_URL + "/dashboard", (req, res, ctx) => {
    // If authenticated, return a mocked user details
    return res(
      ctx.status(200),
      ctx.json({
        siteSize: 57,
        articleSize: 6810,
      })
    );
  }),

  rest.get(BACKEND_ADMIN_API_URL + "/articles", (req, res, ctx) => {
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

  rest.put(BACKEND_ADMIN_API_URL + "/sites/:id", async (req, res, ctx) => {
    const { id } = req.params;
    return res(
      ctx.status(200),
      ctx.json({
        id,
        name: "non-no/fashion",
        url: "https://nonno.hpplus.jp/fashion",
        active: true,
        created_at: "2022-10-23T16:29:46+09:00",
        updated_at: "2022-12-12T10:54:32+09:00",
        edges: {
          site_crawl_rule: {
            id: 1,
            article_selector: "article",
            title_selector: "h1",
            link_selector: "a",
            has_data_to_list: true,
            date_selector: "time",
            date_layout: "YYYY-MM-DD",
            is_time_humanize: false,
            is_spa: false,
            created_at: "2022-10-23T16:29:46+09:00",
            updated_at: "2022-12-12T10:54:32+09:00",
            edges: {},
          },
        },
      })
    );
  }),

  rest.get(
    BACKEND_ADMIN_API_URL + "/sites/get-rss-url/:id",
    (req, res, ctx) => {
      return res(ctx.status(200), ctx.json({ url: "https://c/d/feed" }));
    }
  ),

  rest.get(BACKEND_ADMIN_API_URL + "/site-categories", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        totalCount: 1,
        totalPage: 1,
        pageSize: 100,
        data: [{ id: 17, label: "メディア", sites_count: 3 }],
      })
    );
  }),

  rest.put(
    BACKEND_ADMIN_API_URL + "/site-categories/:id",
    async (req, res, ctx) => {
      const { id } = req.params;
      const { label } = await req.json();
      return res(
        ctx.status(200),
        ctx.json({
          id: id,
          created_at: "2022-11-18T13:25:23+09:00",
          updated_at: "2022-12-08T12:41:12+09:00",
          label: label,
          edges: {},
        })
      );
    }
  ),

  rest.delete(
    BACKEND_ADMIN_API_URL + "/site-categories/:id",
    async (req, res, ctx) => {
      const { id } = req.params;
      return res(ctx.status(200), ctx.json({ message: "deleted" }));
    }
  ),

  rest.post(
    BACKEND_ADMIN_API_URL + "/site-categories",
    async (req, res, ctx) => {
      const { label } = await req.json();
      return res(
        ctx.status(200),
        ctx.json({
          id: 2,
          created_at: "2022-11-18T13:25:23+09:00",
          updated_at: "2022-12-08T12:41:12+09:00",
          label: label,
          edges: {},
        })
      );
    }
  ),

  rest.get(
    BACKEND_ADMIN_API_URL + "/sites/dry-run-crawling/:id",
    (req, res, ctx) => {
      // If authenticated, return a mocked user details
      return res(
        ctx.status(200),
        ctx.json({
          contents:
            '\u003c?xml version="1.0" encoding="UTF-8"?\u003e\u003crss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/"\u003e\n  \u003cchannel\u003e\n    \u003ctitle\u003eファッション – CLASSY.[クラッシィ]\u003c/title\u003e\n    \u003clink\u003ehttps://classy-online.jp/fashion/\u003c/link\u003e\n    \u003cdescription\u003e「FASHION」の記事一覧ページです。【CLASSY.ONLINE】では、アラサー世代が今知りたい最新ニュースやリアルに使えるオシャレのテクニックをファッション、ヘアメーク、ライフスタイル…ジャンルを問わずにまとめて紹介しています！\u003c/description\u003e\n    \u003cpubDate\u003eFri, 09 Dec 2022 12:44:01 +0900\u003c/pubDate\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eコンサバ大人ライターが密かに欲しがっている…！「気になる小物」３…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248822/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eFri, 09 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e【ビームス】2023年の干支「兎」がモチーフ！女の子の憧れfam…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/249092/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eFri, 09 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e超定番アイテム「黒タートルニット」をシンプルだけど、今年っぽく着…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248497/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eFri, 09 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eチェックジャケットにVネックワンピできれいめ冬コーデ【着回しDi…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/246921/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eFri, 09 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e【今日の服装】失敗しない「コート」のおすすめは？【アラサー女子】\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248361/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eFri, 09 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eコンサバ好き大人ライター３名が厳選！「この冬欲しい人気ブランドの…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248772/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e10年、20年先も着られる！名品級の黒コート【伝説のパールチェス…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248493/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e大人女子のあるあるNGコーデ図鑑「制服に見える、チェックパンツ」\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/240070/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e「ニットビーニー」懐かしアイテムが最近キテる！おすすめ大人ニット…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/247376/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eその服、実はちょっと古いです「大きめのほうが今っぽかった…ビッグ…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248237/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e【マイケル･コース】冬のゲレンデを彩る「ellesse」コラボの…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248003/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e黒なのにシンプルすぎない！ベテランスタイリストがおすすめする「大…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248273/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e【G-SHOCK】カレとお揃いで着けたい！メタル×ホワイトの冬限…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/247985/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eホリデーシーズンも毎日活躍！大人気バッグ・EMPORIO ARM…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/246089/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e読者がリアルに「５年以上愛用している、ずっと使える服」２つ\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248766/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003eスウェットやスニーカーと相性がかなりよい！「トレンドが戻ってきた…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248399/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e黒の定番コートでもインナーでは華やかに変化を与えて【着回しDia…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/246916/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n    \u003citem\u003e\n      \u003ctitle\u003e【今日の服装】コートの中に着るだけでこなれる「おすすめアイテム」…\u003c/title\u003e\n      \u003clink\u003ehttps://classy-online.jp/fashion/248360/\u003c/link\u003e\n      \u003cdescription\u003e\u003c/description\u003e\n      \u003cpubDate\u003eThu, 08 Dec 2022 00:00:00 +0900\u003c/pubDate\u003e\n    \u003c/item\u003e\n  \u003c/channel\u003e\n\u003c/rss\u003e',
          count: 18,
        })
      );
    }
  ),
];
