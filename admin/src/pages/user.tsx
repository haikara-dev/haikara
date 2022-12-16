import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import EditUserRoleFormDialog from "@/components/user/EditUserRoleFormDialog";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";
import { useGetUsersQuery } from "@/services/adminApi";
import { User } from "@/features/auth/authSlice";
import PaginationHeader from "@/components/ui/PaginationHeader";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Button from "@mui/material/Button";

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
  };

  const onClickEditButtonHandler = (
    user: User,
    e: React.MouseEvent<HTMLButtonElement>
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
          <PaginationHeader
            totalCount={users.totalCount}
            page={page}
            pageSize={users.pageSize}
            dataSize={users.data.length}
          />
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Email</TableCell>
                  <TableCell>Role</TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {users.data.map((user) => {
                  return (
                    <TableRow key={user.id}>
                      <TableCell>{user.id}</TableCell>
                      <TableCell width={"100%"}>{user.email}</TableCell>
                      <TableCell>{user.role}</TableCell>
                      <TableCell>
                        <Button
                          onClick={onClickEditButtonHandler.bind(this, user)}
                        >
                          Edit
                        </Button>
                      </TableCell>
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          </TableContainer>

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
