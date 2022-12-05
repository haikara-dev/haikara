import AppBar from "@mui/material/AppBar";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Link from "next/link";
import { getAuth } from "firebase/auth";

import { FC, MouseEvent } from "react";
import { useAppDispatch, useAppSelector } from "@/app/hooks";
import { logout, selectAuthUser } from "@/features/auth/authSlice";
import { useRouter } from "next/router";

export type DefaultHeaderProps = {};

const DefaultHeader: FC<DefaultHeaderProps> = () => {
  const auth = getAuth();
  const authUser = useAppSelector(selectAuthUser);
  const dispatch = useAppDispatch();
  const router = useRouter();

  const handleOnClickLogout = async (e: MouseEvent<HTMLButtonElement>) => {
    try {
      await auth.signOut();
      dispatch(logout());
      router.push("/");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <AppBar position="fixed">
      <Toolbar>
        <Link href="/" passHref>
          <Typography
            variant="h6"
            component="a"
            sx={{
              color: "#FFFFFF",
              flexGrow: 1,
              textDecoration: "none",
            }}
          >
            haikara
          </Typography>
        </Link>

        {authUser ? (
          <>
            <Link href="/dashboard" passHref>
              <Button color="inherit">コンソール</Button>
            </Link>

            <Button color="inherit" onClick={handleOnClickLogout}>
              ログアウト
            </Button>
          </>
        ) : (
          <>
            <Link href="/login" passHref>
              <Button color="inherit">ログイン</Button>
            </Link>
            <Link href="/register" passHref>
              <Button color="inherit">登録</Button>
            </Link>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};
export default DefaultHeader;
