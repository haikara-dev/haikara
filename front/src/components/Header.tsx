import {
  AppBar as MuiAppBar,
  AppBarProps as MuiAppBarProps,
  Button,
  Toolbar,
  Typography,
} from "@mui/material";
import Link from "next/link";
import { getAuth } from "firebase/auth";
import { useAuthUserContext } from "@/lib/AuthUser";

import { styled } from "@mui/material/styles";
import { FC } from "react";

const drawerWidth = 240;

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== "open",
})<AppBarProps>(({ theme, open }) => ({
  transition: theme.transitions.create(["margin", "width"], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: `${drawerWidth}px`,
    transition: theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

export type HeaderProps = {
  open?: boolean;
};

const Header: FC<HeaderProps> = ({ open }) => {
  const auth = getAuth();
  const { isAdmin, currentUser, authUser, logout } = useAuthUserContext();
  return (
    <AppBar position="fixed" open={open}>
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
