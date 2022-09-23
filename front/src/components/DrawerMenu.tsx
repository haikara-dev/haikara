import {
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemText,
} from "@mui/material";
import { FC } from "react";
import Link from "next/link";
import { styled } from "@mui/material/styles";
import { useAuthUserContext } from "@/lib/AuthUser";

const drawerWidth = 240;

export type DrawerMenuProps = {
  open?: boolean;
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
  const { isAdmin, authUser } = useAuthUserContext();

  return (
    <Drawer
      variant="permanent"
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
          <List>
            <ListItem disablePadding>
              <Link href="/dashboard" passHref>
                <ListItemButton>
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
          <List>
            <ListItem disablePadding>
              <Link href="/article" passHref>
                <ListItemButton>
                  <ListItemText primary="Article" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/site" passHref>
                <ListItemButton>
                  <ListItemText primary="Site" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/feed" passHref>
                <ListItemButton>
                  <ListItemText primary="Feed" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/user" passHref>
                <ListItemButton>
                  <ListItemText primary="User" />
                </ListItemButton>
              </Link>
            </ListItem>

            <ListItem disablePadding>
              <Link href="/setting" passHref>
                <ListItemButton>
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
