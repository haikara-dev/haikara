import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import { User, useAuthUserContext } from "@/lib/AuthUser";
import EditUserRoleFormDialog from "@/components/user/EditUserRoleFormDialog";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const Users: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );
  const [totalCount, setTotalCount] = useState<number>(0);
  const [totalPage, setTotalPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);

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
      const queryParams = new URLSearchParams({ page: page.toString() });
      const res = await fetch(BACKEND_ADMIN_API_URL + "/users?" + queryParams, {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setTotalCount(json.totalCount);
      setTotalPage(json.totalPage);
      setPageSize(json.pageSize);
      setData(json.data);
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

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { page: page } });
  };

  useEffect(() => {
    setLoading(true);
  }, []);

  useEffect(() => {
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

  useEffect(() => {
    loadData();
  }, [page]);

  return (
    <div>
      <Typography variant="h3" component="h1">
        Users
      </Typography>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <Stack>
            {totalCount}件中　{(page - 1) * pageSize + 1} -{" "}
            {(page - 1) * pageSize + data.length}件
          </Stack>
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
          <Pagination
            page={page}
            count={totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}

      {editTarget && (
        <EditUserRoleFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          user={editTarget}
          updateUserRole={updateUserRole}
          onEndEdit={handleEditClose}
        />
      )}
    </div>
  );
};

Users.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Users;
