import { AppBar, Button, Toolbar, Typography } from "@mui/material";
import Link from "next/link";
import { getAuth } from "firebase/auth";
import { useAuthUserContext } from "@/lib/AuthUser";

const Header = () => {
  const auth = getAuth();
  const { isAdmin, currentUser, authUser, logout } = useAuthUserContext();
  return (
    <AppBar component="header" position="sticky">
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
            {isAdmin && (
              <>
                <Link href="/article" passHref>
                  <Button color="inherit">Article</Button>
                </Link>
                <Link href="/site" passHref>
                  <Button color="inherit">Site</Button>
                </Link>
                <Link href="/feed" passHref>
                  <Button color="inherit">Feed</Button>
                </Link>
                <Link href="/user" passHref>
                  <Button color="inherit">User</Button>
                </Link>
              </>
            )}

            <Typography variant="h6" component="div">
              {currentUser?.email}
            </Typography>
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
