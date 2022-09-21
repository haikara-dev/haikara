import { Alert, Box, Button, Card, TextField, Typography } from "@mui/material";
import {
  getAuth,
  createUserWithEmailAndPassword,
  UserCredential,
} from "firebase/auth";

import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import Link from "next/link";
import React, { ReactElement, useState } from "react";
import { useAuthUserContext } from "@/lib/AuthUser";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "@/pages/_app";
import AuthLayout from "@/components/layouts/AuthLayout";

type FormInput = {
  email: string;
  "new-password": string;
  "confirm-password": string;
};

const schema = yup.object({
  email: yup
    .string()
    .required("必須です")
    .email("正しいメールアドレスを入力してください"),
  "new-password": yup.string().required("必須です"),
  "confirm-password": yup
    .string()
    .oneOf([yup.ref("new-password"), null], "パスワードが一致しません"),
});

const Register: NextPageWithLayout = () => {
  const auth = getAuth();
  const { login } = useAuthUserContext();
  const router = useRouter();
  const [serverError, setServerError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInput>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    setServerError(null);
    try {
      const userCredential: UserCredential =
        await createUserWithEmailAndPassword(
          auth,
          data.email.trim(),
          data["new-password"].trim()
        );
      login(userCredential.user, () => {
        router.push("/");
      });
    } catch (err) {
      if (err instanceof Error) {
        setServerError(err.message);
      }
      console.log(err);
    }
  };

  return (
    <div>
      <Typography variant="h3" component="h1">
        Register
      </Typography>
      <Box
        component="form"
        onSubmit={handleSubmit(onSubmit)}
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 2,
          mt: 3,
        }}
      >
        {serverError && <Alert severity="error">{serverError}</Alert>}
        <TextField
          label="email"
          placeholder="email"
          autoFocus={true}
          {...register("email")}
          error={errors.email ? true : false}
          helperText={errors.email?.message}
          variant="outlined"
        />

        <TextField
          label="password"
          type="password"
          placeholder="password"
          {...register("new-password")}
          error={errors["new-password"] ? true : false}
          helperText={errors["new-password"]?.message}
          variant="outlined"
        />

        <TextField
          label="confirm password"
          type="password"
          placeholder="password"
          {...register("confirm-password")}
          error={errors["confirm-password"] ? true : false}
          helperText={errors["confirm-password"]?.message}
          variant="outlined"
        />

        <Button type="submit" aria-label="register" variant="contained">
          Register
        </Button>
        <Link href="/login" passHref>
          <Button aria-label="login">Login</Button>
        </Link>
      </Box>
    </div>
  );
};

Register.getLayout = (page: ReactElement) => <AuthLayout>{page}</AuthLayout>;

export default Register;
