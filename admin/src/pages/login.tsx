import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import { User as AuthUser } from "@firebase/auth";
import {
  getAuth,
  signInWithEmailAndPassword,
  UserCredential,
} from "firebase/auth";
import { SubmitHandler, useForm } from "react-hook-form";
import React, { ReactElement, useEffect, useState } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import Link from "next/link";
import { useRouter } from "next/router";
import AuthLayout from "@/components/layouts/AuthLayout";
import { NextPageWithLayout } from "@/pages/_app";
import { useAppDispatch, useAppSelector } from "@/app/hooks";
import {
  login,
  selectAuthUser,
  setCurrentUser,
} from "@/features/auth/authSlice";
import {
  useCreateUserMutation,
  useLazyGetCurrentUserQuery,
  userApi,
} from "@/services/userApi";

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
  const authUser = useAppSelector(selectAuthUser);
  const router = useRouter();
  const dispatch = useAppDispatch();

  const [serverError, setServerError] = useState<string | null>(null);
  const [getCurrentUser] = useLazyGetCurrentUserQuery();
  const [createUser] = useCreateUserMutation();

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

      dispatch(login(userCredential.user));
    } catch (err) {
      if (err instanceof Error) {
        setServerError(err.message);
      }
      console.log(err);
    }
  };

  useEffect(() => {
    const onLogin = async (authUser: AuthUser) => {
      try {
        await createUser(authUser).unwrap();

        const currentUser = await getCurrentUser().unwrap();

        if (currentUser) {
          dispatch(setCurrentUser(currentUser));
          await router.push("/dashboard");
        }
      } catch (err) {
        if (err instanceof Error) {
          setServerError(err.message);
        }
        console.log(err);
      }
    };
    if (authUser) {
      onLogin(authUser);
    }
  }, [authUser]);

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
