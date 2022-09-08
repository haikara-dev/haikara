import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React from "react";
import {
  Box,
  TextField,
  Button,
  DialogTitle,
  DialogContent,
  Dialog,
  DialogContentText,
  DialogActions,
} from "@mui/material";

import { Site } from "@/pages/site";

type FormInput = {
  name: string;
  url: string;
  feed_url: string;
};

const schema = yup.object({
  name: yup.string().required("必須です"),
  url: yup.string().required("必須です").url("正しいURLを入力してください"),
  feed_url: yup
    .string()
    .required("必須です")
    .url("正しいURLを入力してください"),
});

export type AddSiteFormProps = {
  open: boolean;
  handleClose: () => void;
  site: Site;
  updateSite: (
    id: number,
    name: string,
    url: string,
    feed_url: string,
    active: boolean
  ) => void;
  getRssUrl: (id: number) => Promise<string>;
  onEndEdit: () => void;
};

const EditSiteFormDialog: React.FC<AddSiteFormProps> = ({
  open,
  handleClose,
  site,
  updateSite,
  getRssUrl,
  onEndEdit,
}) => {
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
      if (trimmedfeed_url.length === 0) return;

      await updateSite(site.id, trimmedName, trimmedUrl, trimmedfeed_url, true);

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

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth={true}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle>Edit Site</DialogTitle>
        <DialogContent>
          <DialogContentText>Siteの名前とURLを編集できます。</DialogContentText>
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
          </Box>
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
