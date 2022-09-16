import type { NextPage } from "next";
import Head from "next/head";
import { Container, Typography, Box, Stack, Card, Button } from "@mui/material";
import React, { useEffect, useState } from "react";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { User, useAuthUserContext } from "@/lib/AuthUser";
import EditUserRoleFormDialog from "@/components/user/EditUserRoleFormDialog";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const Users: NextPage = () => {
  const [editOpen, setEditOpen] = useState(false);
  const [editTarget, setEditTarget] = useState<User | null>(null);

  const [data, setData] = useState<User[]>([]);
  const [isLoading, setLoading] = useState(false);
  const { authUser } = useAuthUserContext();

  const handleEditOpen = (user: User) => {
    setEditOpen(true);
    setEditTarget(user);
  };

  const handleEditClose = () => {
    setEditOpen(false);
    setEditTarget(null);
  };

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/users", {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setData(json);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const updateUserRole = async (id: number, role: string) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/users/role/"),
        {
          method: "PATCH",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            role: role,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const onClickTextHandler = (
    user: User,
    e: React.MouseEvent<HTMLDivElement>
  ) => {
    e.preventDefault();
    handleEditOpen(user);
  };

  useEffect(() => {
    setLoading(true);
    loadData();
  }, []);

  return (
    <div>
      <Head>
        <title>haikara</title>
        <meta name="description" content="haikara" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <Box
        component="main"
        sx={{
          minHeight: "100vh",
        }}
      >
        <Container
          sx={{
            p: 2,
          }}
        >
          <Typography variant="h3" component="h1">
            Users
          </Typography>

          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <Stack gap={2} mt={2} pr={8}>
              {data.map((user) => {
                return (
                  <Card key={user.id}>
                    <Stack direction="row" gap={3} alignItems="center">
                      <div>{user.id}</div>
                      <div onClick={onClickTextHandler.bind(this, user)}>
                        {user.email}
                      </div>
                      <div>{user.role}</div>
                    </Stack>
                  </Card>
                );
              })}
            </Stack>
          )}
        </Container>
      </Box>

      {editTarget && (
        <EditUserRoleFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          user={editTarget}
          updateUserRole={updateUserRole}
          onEndEdit={handleEditClose}
        />
      )}

      <Footer />
    </div>
  );
};

export default Users;
