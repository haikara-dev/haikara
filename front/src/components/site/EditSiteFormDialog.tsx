import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React, { useEffect } from "react";

import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import Dialog from "@mui/material/Dialog";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import FormHelperText from "@mui/material/FormHelperText";
import FormControl from "@mui/material/FormControl";

import { Site, SiteWithSiteCrawlRule } from "@/pages/site";

type FormInput = {
  name: string;
  url: string;
  feed_url: string;
  cannot_crawl: boolean;
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
  cannot_crawl: yup.boolean(),

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

export type AddSiteFormProps = {
  open: boolean;
  handleClose: () => void;
  site: Site;
  updateSite: (site: SiteWithSiteCrawlRule) => void;
  getRssUrl: (id: number) => Promise<string>;
  onEndEdit: () => void;
  loadSiteWithSiteCrawlRule: (
    id: number
  ) => Promise<SiteWithSiteCrawlRule | null>;
};

const EditSiteFormDialog: React.FC<AddSiteFormProps> = ({
  open,
  handleClose,
  site,
  updateSite,
  getRssUrl,
  onEndEdit,
  loadSiteWithSiteCrawlRule,
}) => {
  const [siteWithSiteCrawlRule, setSiteWithSiteCrawlRule] =
    React.useState<SiteWithSiteCrawlRule | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<FormInput>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedName = data.name.trim();
      if (trimmedName.length === 0) return;

      const trimmedUrl = data.url.trim();
      if (trimmedUrl.length === 0) return;

      const trimmedfeed_url = data.feed_url.trim();

      await updateSite({
        id: site.id,
        name: trimmedName,
        url: trimmedUrl,
        feed_url: trimmedfeed_url,
        active: site.active,
        cannot_crawl_at: site.cannot_crawl_at,
        cannot_crawl: data.cannot_crawl,
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
      });

      onEndEdit();
    } catch (err) {
      console.log(err);
    }
  };

  const onClickGetFeedUrlHandler = async (
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    const url = await getRssUrl(site.id);
    if (url !== "") {
      setValue("feed_url", url, { shouldValidate: true });
    }
  };

  useEffect(() => {
    const loadSite = async () => {
      const siteWithSiteCrawlRule = await loadSiteWithSiteCrawlRule(site.id);
      setSiteWithSiteCrawlRule(siteWithSiteCrawlRule);
    };
    if (site === null) {
      setSiteWithSiteCrawlRule(null);
    } else {
      loadSite();
    }
  }, [site]);

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth={true}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle>Edit Site</DialogTitle>
        <DialogContent>
          <DialogContentText>Siteの名前とURLを編集できます。</DialogContentText>

          {siteWithSiteCrawlRule ? (
            <Box display="flex" flexDirection="column" gap={2}>
              <TextField
                label="Site Name"
                autoFocus={true}
                error={errors.name ? true : false}
                helperText={errors.name && errors.name.message}
                variant="outlined"
                placeholder="name"
                defaultValue={site.name}
                {...register("name")}
                sx={{ flexGrow: 1 }}
              />

              <TextField
                label="Site URL"
                error={errors.url ? true : false}
                helperText={errors.url && errors.url.message}
                variant="outlined"
                placeholder="url"
                defaultValue={site.url}
                {...register("url")}
                sx={{ flexGrow: 1 }}
              />

              <Button onClick={onClickGetFeedUrlHandler}>Get Feed URL</Button>

              <TextField
                label="Feed URL"
                error={errors.feed_url ? true : false}
                helperText={errors.feed_url && errors.feed_url.message}
                variant="outlined"
                placeholder="Feed URL"
                defaultValue={site.feed_url}
                {...register("feed_url")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />

              <div>Site Crawl Rule</div>

              <TextField
                label="article_selector"
                error={errors.article_selector ? true : false}
                helperText={
                  errors.article_selector && errors.article_selector.message
                }
                variant="outlined"
                placeholder="article_selector"
                defaultValue={
                  siteWithSiteCrawlRule.site_crawl_rule.article_selector
                }
                {...register("article_selector")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />

              <TextField
                label="title_selector"
                error={errors.title_selector ? true : false}
                helperText={
                  errors.title_selector && errors.title_selector.message
                }
                variant="outlined"
                placeholder="title_selector"
                defaultValue={
                  siteWithSiteCrawlRule.site_crawl_rule.title_selector
                }
                {...register("title_selector")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />
              <TextField
                label="link_selector"
                error={errors.link_selector ? true : false}
                helperText={
                  errors.link_selector && errors.link_selector.message
                }
                variant="outlined"
                placeholder="link_selector"
                defaultValue={
                  siteWithSiteCrawlRule.site_crawl_rule.link_selector
                }
                {...register("link_selector")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />
              <TextField
                label="description_selector"
                error={errors.description_selector ? true : false}
                helperText={
                  errors.description_selector &&
                  errors.description_selector.message
                }
                variant="outlined"
                placeholder="description_selector"
                defaultValue={
                  siteWithSiteCrawlRule.site_crawl_rule.description_selector
                }
                {...register("description_selector")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />

              <FormControl error={errors.has_data_to_list ? true : false}>
                <FormControlLabel
                  control={
                    <Checkbox
                      defaultChecked={
                        siteWithSiteCrawlRule.site_crawl_rule.has_data_to_list
                      }
                      {...register("has_data_to_list")}
                    />
                  }
                  label="has_data_to_list"
                />
                {errors.has_data_to_list && (
                  <FormHelperText>
                    {errors.has_data_to_list.message}
                  </FormHelperText>
                )}
              </FormControl>

              <TextField
                label="date_selector"
                error={errors.date_selector ? true : false}
                helperText={
                  errors.date_selector && errors.date_selector.message
                }
                variant="outlined"
                placeholder="date_selector"
                defaultValue={
                  siteWithSiteCrawlRule.site_crawl_rule.date_selector
                }
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
                defaultValue={siteWithSiteCrawlRule.site_crawl_rule.date_layout}
                {...register("date_layout")}
                sx={{ flexGrow: 1 }}
                InputLabelProps={{ shrink: true }}
              />
              <FormControl error={errors.is_time_humanize ? true : false}>
                <FormControlLabel
                  control={
                    <Checkbox
                      defaultChecked={
                        siteWithSiteCrawlRule.site_crawl_rule.is_time_humanize
                      }
                      {...register("is_time_humanize")}
                    />
                  }
                  label="is_time_humanize"
                />
                {errors.is_time_humanize && (
                  <FormHelperText>
                    {errors.is_time_humanize.message}
                  </FormHelperText>
                )}
              </FormControl>

              <FormControl error={errors.is_spa ? true : false}>
                <FormControlLabel
                  control={
                    <Checkbox
                      defaultChecked={
                        siteWithSiteCrawlRule.site_crawl_rule.is_spa
                      }
                      {...register("is_spa")}
                    />
                  }
                  label="is_spa"
                />
                {errors.is_spa && (
                  <FormHelperText>{errors.is_spa.message}</FormHelperText>
                )}
              </FormControl>

              <FormControl error={errors.cannot_crawl ? true : false}>
                <FormControlLabel
                  control={
                    <Checkbox
                      defaultChecked={site.cannot_crawl}
                      {...register("cannot_crawl")}
                    />
                  }
                  label={
                    site.cannot_crawl
                      ? `cannot_crawl( ${new Date(
                          site.cannot_crawl_at
                        ).toLocaleString()} )`
                      : "cannot_crawl"
                  }
                />
                {errors.cannot_crawl && (
                  <FormHelperText>{errors.cannot_crawl.message}</FormHelperText>
                )}
              </FormControl>
            </Box>
          ) : (
            <div>loading...</div>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button type="submit">Save</Button>
        </DialogActions>
      </Box>
    </Dialog>
  );
};

export default EditSiteFormDialog;
