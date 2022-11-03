import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React from "react";
import {
  Box,
  Button,
  DialogTitle,
  DialogContent,
  Dialog,
  DialogContentText,
  DialogActions,
  TextField,
} from "@mui/material";
import {
  SiteCategory,
  useUpdateSiteCategoryMutation,
} from "@/services/adminApi";

type FormInput = {
  label: string;
};

const schema = yup.object({
  label: yup.string().required("必須です"),
});

export type EditSiteCategoryFormDialogProps = {
  open: boolean;
  handleClose: () => void;
  siteCategory: SiteCategory;
};

const EditSiteCategoryFormDialog: React.FC<EditSiteCategoryFormDialogProps> = ({
  open,
  handleClose,
  siteCategory,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInput>({
    defaultValues: {
      label: siteCategory.label,
    },
    resolver: yupResolver(schema),
  });

  const [updateSiteCategory] = useUpdateSiteCategoryMutation();

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedLabel = data.label.trim();
      if (trimmedLabel.length === 0) return;

      await updateSiteCategory({
        id: siteCategory.id,
        body: {
          label: trimmedLabel,
        },
      });

      handleClose();
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth={true}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle>Edit Site Category</DialogTitle>
        <DialogContent>
          <DialogContentText>入力してください。</DialogContentText>
          <Box display="flex" flexDirection="column" gap={2}>
            <TextField
              label="Site Category Label"
              autoFocus={true}
              error={errors.label ? true : false}
              helperText={errors.label && errors.label.message}
              variant="outlined"
              placeholder="label"
              {...register("label")}
              sx={{ flexGrow: 1 }}
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

export default EditSiteCategoryFormDialog;
