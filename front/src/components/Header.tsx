import AppBar from "@mui/material/AppBar";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Link from "next/link";
import { getAuth } from "firebase/auth";
import { useAuthUserContext } from "@/lib/AuthUser";

import { FC } from "react";

export type HeaderProps = {
  handleToggleDrawer?: () => void;
};

const Header: FC<HeaderProps> = ({ handleToggleDrawer }) => {
  const auth = getAuth();
  const { authUser, logout } = useAuthUserContext();
  return (
    <AppBar
      position="fixed"
      sx={{
        zIndex: (theme) => theme.zIndex.drawer + 1,
      }}
    >
      <Toolbar>
        <IconButton
          color="inherit"
          aria-label="open drawer"
          onClick={handleToggleDrawer}
          edge="start"
          sx={{ mr: 2 }}
        >
          <MenuIcon />
        </IconButton>
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

            {/*<Typography variant="h6" component="div">*/}
            {/*  {currentUser?.email}*/}
            {/*</Typography>*/}
            <Button
              color="inherit"
              onClick={async () => {
                try {
                  await auth.signOut();
                  logout(() => {});
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
export default Header;
