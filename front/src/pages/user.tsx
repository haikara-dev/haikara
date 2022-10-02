import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import EditUserRoleFormDialog from "@/components/user/EditUserRoleFormDialog";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";
import { useGetUsersQuery } from "@/services/adminApi";
import { User } from "@/features/auth/authSlice";

const Users: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );

  const {
    data: users = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
    refetch,
  } = useGetUsersQuery(page);

  const [editOpen, setEditOpen] = useState(false);
  const [editTarget, setEditTarget] = useState<User | null>(null);

  const handleEditOpen = (user: User) => {
    setEditOpen(true);
    setEditTarget(user);
  };

  const handleEditClose = () => {
    setEditOpen(false);
    setEditTarget(null);
  };

  const handleChangedRole = () => {
    handleEditClose();
    refetch();
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
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

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
            {users.totalCount}件中　{(page - 1) * users.pageSize + 1} -{" "}
            {(page - 1) * users.pageSize + users.data.length}件
          </Stack>
          <Stack gap={2} mt={2} pr={8}>
            {users.data.map((user) => {
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
            count={users.totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}

      {editTarget && (
        <EditUserRoleFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          user={editTarget}
          onEndEdit={handleChangedRole}
        />
      )}
    </div>
  );
};

Users.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Users;
