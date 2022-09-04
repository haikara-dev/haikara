import { useForm, SubmitHandler, SubmitErrorHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import React from "react";
import { Box, TextField, FormHelperText } from "@mui/material";

import { Todo } from "@/pages/todo";
type FormInput = {
  text: string;
};

const schema = yup.object({
  text: yup.string().required("必須です"),
});

export type AddTodoFormProps = {
  todo: Todo;
  updateTodo: (id: number, text: string) => void;
  onEndEdit: () => void;
};

const EditTodoForm: React.FC<AddTodoFormProps> = ({
  todo,
  updateTodo,
  onEndEdit,
}) => {
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
      await updateTodo(todo.id, trimmedText);
    } catch (err) {
      console.log(err);
    }
    onEndEdit();
  };

  const onBlur: SubmitHandler<FormInput> = async (data) => {
    try {
      const trimmedText = data.text.trim();
      if (trimmedText.length === 0) {
        reset();
        onEndEdit();
      } else {
        await updateTodo(todo.id, trimmedText);
        onEndEdit();
      }
    } catch (err) {
      console.log(err);
    }
  };

  const onError: SubmitErrorHandler<FormInput> = async (err) => {
    onEndEdit();
  };

  return (
    <Box component="form" onSubmit={handleSubmit(onSubmit)}>
      <Box display="flex" gap={1}>
        <TextField
          autoFocus={true}
          error={errors.text ? true : false}
          variant="outlined"
          placeholder="input here!"
          defaultValue={todo.text}
          {...register("text")}
          sx={{ flexGrow: 1 }}
          onBlur={handleSubmit(onBlur, onError)}
        />
      </Box>
      <FormHelperText error={errors.text ? true : false}>
        {errors.text && errors.text.message}
      </FormHelperText>
    </Box>
  );
};

export default EditTodoForm;
