import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React, { useEffect } from "react";

import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import FormHelperText from "@mui/material/FormHelperText";
import FormControl from "@mui/material/FormControl";

import Stack from "@mui/material/Stack";
import { useRouter } from "next/router";
import {
  useAddSiteMutation,
  useGetSiteCategoriesQuery,
  useGetSiteRssUrlByUrlMutation,
} from "@/services/adminApi";
import { FormGroup } from "@mui/material";

type FormInput = {
  name: string;
  url: string;
  feed_url: string;
  site_category_ids: number[] | boolean;
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

const schema = yup.object({
  name: yup.string().required("必須です"),
  url: yup.string().required("必須です").url("正しいURLを入力してください"),
  feed_url: yup.string().url("正しいURLを入力してください"),

  article_selector: yup.string(),
  title_selector: yup.string(),
  link_selector: yup.string(),
  description_selector: yup.string(),
  has_data_to_list: yup.boolean(),
  date_selector: yup.string(),
  date_layout: yup.string(),
  is_time_humanize: yup.boolean(),
  is_spa: yup.boolean(),
});

const AddSiteForm = () => {
  const router = useRouter();

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
    setValue,
    getValues,
  } = useForm<FormInput>({
    defaultValues: {
      name: "",
      url: "",
      feed_url: "",

      article_selector: "",
      title_selector: "",
      link_selector: "",
      description_selector: "",
      has_data_to_list: true,
      date_selector: "",
      date_layout: "",
      is_time_humanize: false,
      is_spa: false,
    },
    resolver: yupResolver(schema),
  });

  const [addSite] = useAddSiteMutation();
  const [getSiteRssUrlByUrl] = useGetSiteRssUrlByUrlMutation();
  const { data: siteCategories } = useGetSiteCategoriesQuery({ page: 1 });

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedName = data.name.trim();
      if (trimmedName.length === 0) return;

      const trimmedUrl = data.url.trim();
      if (trimmedUrl.length === 0) return;

      const trimmedFeedUrl = data.feed_url.trim();

      const formatSiteCategoryIds = () => {
        if (data.site_category_ids === false) {
          return [];
        }

        if (Array.isArray(data.site_category_ids)) {
          return data.site_category_ids.map((val) => {
            return Number(val);
          });
        } else {
          return [Number(data.site_category_ids)];
        }
      };

      const siteCategoryIds = formatSiteCategoryIds();

      await addSite({
        body: {
          // TODO: ここでidを生成しているが、使わない
          id: 0,
          name: trimmedName,
          url: trimmedUrl,
          feed_url: trimmedFeedUrl,
          active: false,
          cannot_crawl: false,
          cannot_crawl_at: "",
          site_crawl_rule: {
            article_selector: data.article_selector,
            title_selector: data.title_selector,
            link_selector: data.link_selector,
            description_selector: data.description_selector,
            has_data_to_list: data.has_data_to_list,
            date_selector: data.date_selector,
            date_layout: data.date_layout,
            is_time_humanize: data.is_time_humanize,
            is_spa: data.is_spa,
          },
          site_category_ids: siteCategoryIds,
        },
      });

      reset();
      goListPage();
    } catch (err) {
      console.log(err);
    }
  };

  const onClickGetFeedUrlHandler = async (
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    const res = await getSiteRssUrlByUrl(getValues("url")).unwrap();
    if (res.url !== "") {
      setValue("feed_url", res.url, {
        shouldValidate: true,
      });
    }
  };
  const goListPage = () => {
    router.push("/sites");
  };
  const goBackPage = () => {
    router.back();
  };
  const onClickBackPageHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    goBackPage();
  };

  useEffect(() => {
    reset();
  }, []);

  return (
    <Box component="form" onSubmit={handleSubmit(onSubmit)}>
      <div>Siteの名前とURLを入力してください。</div>

      <Box display="flex" flexDirection="column" gap={2}>
        <TextField
          label="Site Name"
          autoFocus={true}
          error={errors.name ? true : false}
          helperText={errors.name && errors.name.message}
          variant="outlined"
          placeholder="name"
          {...register("name")}
          sx={{ flexGrow: 1 }}
        />

        <TextField
          label="Site URL"
          error={errors.url ? true : false}
          helperText={errors.url && errors.url.message}
          variant="outlined"
          placeholder="url"
          {...register("url")}
          sx={{ flexGrow: 1 }}
        />

        <Button onClick={onClickGetFeedUrlHandler}>Get Feed URL</Button>

        <TextField
          label="Feed URL"
          error={errors.feed_url ? true : false}
          helperText={errors.feed_url && errors.feed_url.message}
          variant="outlined"
          placeholder="Feed Url"
          {...register("feed_url")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />

        <div>Site Category</div>
        {siteCategories && (
          <FormGroup>
            {siteCategories.data.map((siteCategory) => (
              <FormControlLabel
                key={siteCategory.id}
                control={
                  <Checkbox
                    value={siteCategory.id}
                    {...register("site_category_ids")}
                  />
                }
                label={siteCategory.label}
              />
            ))}
          </FormGroup>
        )}

        <div>Site Crawl Rule</div>

        <TextField
          label="article_selector"
          error={errors.article_selector ? true : false}
          helperText={
            errors.article_selector && errors.article_selector.message
          }
          variant="outlined"
          placeholder="article_selector"
          {...register("article_selector")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />

        <TextField
          label="title_selector"
          error={errors.title_selector ? true : false}
          helperText={errors.title_selector && errors.title_selector.message}
          variant="outlined"
          placeholder="title_selector"
          {...register("title_selector")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />
        <TextField
          label="link_selector"
          error={errors.link_selector ? true : false}
          helperText={errors.link_selector && errors.link_selector.message}
          variant="outlined"
          placeholder="link_selector"
          {...register("link_selector")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />
        <TextField
          label="description_selector"
          error={errors.description_selector ? true : false}
          helperText={
            errors.description_selector && errors.description_selector.message
          }
          variant="outlined"
          placeholder="description_selector"
          {...register("description_selector")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />

        <FormControl error={errors.has_data_to_list ? true : false}>
          <FormControlLabel
            control={
              <Checkbox
                {...register("has_data_to_list")}
                defaultChecked={getValues("has_data_to_list")}
              />
            }
            label="has_data_to_list"
          />
          {errors.has_data_to_list && (
            <FormHelperText>{errors.has_data_to_list.message}</FormHelperText>
          )}
        </FormControl>

        <TextField
          label="date_selector"
          error={errors.date_selector ? true : false}
          helperText={errors.date_selector && errors.date_selector.message}
          variant="outlined"
          placeholder="date_selector"
          {...register("date_selector")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />
        <TextField
          label="date_layout"
          error={errors.date_layout ? true : false}
          helperText={errors.date_layout && errors.date_layout.message}
          variant="outlined"
          placeholder="date_layout"
          {...register("date_layout")}
          sx={{ flexGrow: 1 }}
          InputLabelProps={{ shrink: true }}
        />

        <FormControl error={errors.is_time_humanize ? true : false}>
          <FormControlLabel
            control={
              <Checkbox
                {...register("is_time_humanize")}
                defaultChecked={getValues("is_time_humanize")}
              />
            }
            label="is_time_humanize"
          />
          {errors.is_time_humanize && (
            <FormHelperText>{errors.is_time_humanize.message}</FormHelperText>
          )}
        </FormControl>

        <FormControl error={errors.is_spa ? true : false}>
          <FormControlLabel
            control={
              <Checkbox
                {...register("is_spa")}
                defaultChecked={getValues("is_spa")}
              />
            }
            label="is_spa"
          />
          {errors.is_spa && (
            <FormHelperText>{errors.is_spa.message}</FormHelperText>
          )}
        </FormControl>
      </Box>

      <Stack direction="row" spacing={2} sx={{ mt: 2 }} justifyContent="center">
        <Button variant="outlined" onClick={onClickBackPageHandler}>
          戻る
        </Button>
        <Button type="submit" variant="contained">
          Add
        </Button>
      </Stack>
    </Box>
  );
};

export default AddSiteForm;
