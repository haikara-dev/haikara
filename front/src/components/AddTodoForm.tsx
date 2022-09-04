import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React from "react";
import {
  Box,
  TextField,
  Button,
  FormHelperText,
  IconButton,
} from "@mui/material";

import AddCircleIcon from "@mui/icons-material/AddCircle";
type FormInput = {
  text: string;
};

const schema = yup.object({
  text: yup.string().required("必須です"),
});

export type AddTodoFormProps = {
  addTodo: (text: string) => void;
};

const AddTodoForm: React.FC<AddTodoFormProps> = ({ addTodo }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<FormInput>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedText = data.text.trim();
      if (trimmedText.length === 0) return;
      await addTodo(trimmedText);
      reset();
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit(onSubmit)}>
      <Box display="flex" gap={1}>
        <TextField
          error={errors.text ? true : false}
          variant="outlined"
          placeholder="input here!"
          {...register("text")}
          sx={{ flexGrow: 1 }}
        />
        <IconButton type="submit" aria-label="add" size="large">
          <AddCircleIcon fontSize="large" />
        </IconButton>
      </Box>
      <FormHelperText error={errors.text ? true : false}>
        {errors.text && errors.text.message}
      </FormHelperText>
    </Box>
  );
};

export default AddTodoForm;
