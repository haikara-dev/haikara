import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import {
  getAuth,
  signInWithEmailAndPassword,
  UserCredential,
} from "firebase/auth";
import { SubmitHandler, useForm } from "react-hook-form";
import React, { ReactElement, useState } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import Link from "next/link";
import { useAuthUserContext } from "@/lib/AuthUser";
import { useRouter } from "next/router";
import AuthLayout from "@/components/layouts/AuthLayout";
import { NextPageWithLayout } from "@/pages/_app";

type FormInput = {
  email: string;
  password: string;
};

const schema = yup.object({
  email: yup
    .string()
    .required("必須です")
    .email("正しいメールアドレスを入力してください"),
  password: yup.string().required("必須です"),
});

const Login: NextPageWithLayout = () => {
  const auth = getAuth();
  const { authUser, login } = useAuthUserContext();
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
      const userCredential: UserCredential = await signInWithEmailAndPassword(
        auth,
        data.email.trim(),
        data.password.trim()
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
        Login
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
          {...register("password")}
          error={errors.password ? true : false}
          helperText={errors.password?.message}
          variant="outlined"
        />

        <Button type="submit" aria-label="login" variant="contained">
          Login
        </Button>
        <Link href="/register" passHref>
          <Button aria-label="login">Register</Button>
        </Link>
      </Box>
    </div>
  );
};

Login.getLayout = (page: ReactElement) => <AuthLayout>{page}</AuthLayout>;

export default Login;
