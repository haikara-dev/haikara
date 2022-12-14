import { useForm, SubmitHandler, Controller } from "react-hook-form";
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
  FormControl,
  Radio,
  RadioGroup,
  FormControlLabel,
  FormLabel,
  FormHelperText,
} from "@mui/material";
import { User } from "@/features/auth/authSlice";
import { useUpdateUserRoleMutation } from "@/services/adminApi";

type FormInput = {
  role: string;
};

const schema = yup.object({
  role: yup.string().required("必須です"),
});

export type AddSiteFormProps = {
  open: boolean;
  handleClose: () => void;
  user: User;
  onEndEdit: () => void;
};

const EditUserRoleFormDialog: React.FC<AddSiteFormProps> = ({
  open,
  handleClose,
  user,
  onEndEdit,
}) => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInput>({
    resolver: yupResolver(schema),
  });

  const [updateUserRole, result] = useUpdateUserRoleMutation();

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedRole = data.role.trim();
      if (trimmedRole.length === 0) return;

      await updateUserRole({
        id: user.id,
        body: {
          role: trimmedRole,
        },
      });

      onEndEdit();
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth={true}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle>Edit User Role</DialogTitle>
        <DialogContent>
          <DialogContentText>Roleを選択してください。</DialogContentText>
          <Box display="flex" flexDirection="column" gap={2}>
            <FormControl error={errors.role ? true : false}>
              <FormLabel id="role-group">Role</FormLabel>
              {errors.role && (
                <FormHelperText>{errors.role.message}</FormHelperText>
              )}
              <Controller
                name="role"
                control={control}
                defaultValue={user.role}
                render={({ field }) => (
                  <RadioGroup {...field}>
                    <FormControlLabel
                      value="user"
                      control={<Radio />}
                      label="User"
                    />
                    <FormControlLabel
                      value="admin"
                      control={<Radio />}
                      label="Admin"
                    />
                  </RadioGroup>
                )}
              />
            </FormControl>
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

export default EditUserRoleFormDialog;
