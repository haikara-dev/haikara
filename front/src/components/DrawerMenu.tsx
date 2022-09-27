import Divider from "@mui/material/Divider";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import { FC } from "react";
import Link from "next/link";
import styled from "@mui/material/styles/styled";
import { useAuthUserContext } from "@/lib/AuthUser";
import { useRouter } from "next/router";

const drawerWidth = 140;

export type DrawerMenuProps = {
  open: boolean;
};

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

export const DrawerMenu: FC<DrawerMenuProps> = ({ open }) => {
  const router = useRouter();
  const { isAdmin, authUser } = useAuthUserContext();

  return (
    <Drawer
      variant="persistent"
      anchor="left"
      open={open}
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        "& .MuiDrawer-paper": {
          width: drawerWidth,
          boxSizing: "border-box",
        },
      }}
    >
      <DrawerHeader />
      {authUser && (
        <>
          <Divider />
          <List component="nav">
            <ListItem disablePadding>
              <Link href="/dashboard" passHref>
                <ListItemButton selected={router.pathname === "/dashboard"}>
                  <ListItemText primary="Dashboard" />
                </ListItemButton>
              </Link>
            </ListItem>
          </List>
        </>
      )}
      {isAdmin && (
        <>
          <Divider />
          <List component="nav">
            <ListItem disablePadding>
              <Link href="/article" passHref>
                <ListItemButton selected={router.pathname === "/article"}>
                  <ListItemText primary="Article" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/sites" passHref>
                <ListItemButton selected={router.pathname.startsWith("/sites")}>
                  <ListItemText primary="Sites" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/feed" passHref>
                <ListItemButton selected={router.pathname === "/feed"}>
                  <ListItemText primary="Feed" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/user" passHref>
                <ListItemButton selected={router.pathname === "/user"}>
                  <ListItemText primary="User" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/setting" passHref>
                <ListItemButton selected={router.pathname === "/setting"}>
                  <ListItemText primary="Setting" />
                </ListItemButton>
              </Link>
            </ListItem>
          </List>
        </>
      )}
    </Drawer>
  );
};
export default DrawerMenu;
