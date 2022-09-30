import AppBar from "@mui/material/AppBar";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Link from "next/link";
import { getAuth } from "firebase/auth";

import { FC } from "react";
import { useAppDispatch } from "@/app/hooks";
import {
  logout,
  selectAuthUser,
  useAuthSelector,
} from "@/features/auth/authSlice";

export type DefaultHeaderProps = {};

const DefaultHeader: FC<DefaultHeaderProps> = () => {
  const auth = getAuth();
  const authUser = useAuthSelector(selectAuthUser);
  const dispatch = useAppDispatch();
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

            <Button
              color="inherit"
              onClick={async () => {
                try {
                  await auth.signOut();
                  dispatch(logout());
                } catch (error) {
                  console.error(error);
                }
              }}
            >
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
