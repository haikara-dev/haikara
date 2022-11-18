import { NextPageWithLayout } from "@/pages/_app";
import Typography from "@mui/material/Typography";
import React, { ReactElement, useCallback, useEffect, useState } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import Stack from "@mui/material/Stack";
import PaginationHeader from "@/components/ui/PaginationHeader";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import Button from "@mui/material/Button";
import DeleteIcon from "@mui/icons-material/Delete";
import Pagination from "@mui/material/Pagination";
import { useRouter } from "next/router";
import {
  SiteCategory,
  useDeleteSiteCategoryMutation,
  useGetSiteCategoriesQuery,
} from "@/services/adminApi";
import AddSiteCategoryFormDialog from "@/components/site-category/AddSiteCategoryFormDialog";
import EditSiteCategoryFormDialog from "@/components/site-category/EditSiteCategoryFormDialog";
import IconButton from "@mui/material/IconButton";

const SiteCategories: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );

  const {
    data: siteCategories = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
  } = useGetSiteCategoriesQuery({ page });

  const [deleteSiteCategory] = useDeleteSiteCategoryMutation();

  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [editTarget, setEditTarget] = useState<SiteCategory | null>(null);

  const handleAddOpen = () => {
    setAddOpen(true);
  };

  const handleAddClose = () => {
    setAddOpen(false);
  };

  const handleEditOpen = (siteCategory: SiteCategory) => {
    setEditOpen(true);
    setEditTarget(siteCategory);
  };

  const handleEditClose = () => {
    setEditOpen(false);
    setEditTarget(null);
  };

  const onClickAddButtonHandler = useCallback(
    (e: React.MouseEvent<HTMLButtonElement>) => {
      e.preventDefault();
      handleAddOpen();
    },
    []
  );

  const onClickEditButtonHandler = useCallback(
    (siteCategory: SiteCategory, e: React.MouseEvent<HTMLButtonElement>) => {
      e.preventDefault();
      handleEditOpen(siteCategory);
    },
    []
  );

  const onClickDeleteHandler = useCallback(
    async (id: number, e: React.MouseEvent<HTMLButtonElement>) => {
      e.preventDefault();
      await deleteSiteCategory(id);
    },
    []
  );

  const handleChangePagination = useCallback(
    (e: React.ChangeEvent<unknown>, page: number) => {
      router.push({ query: { page: page } });
    },
    [page]
  );

  useEffect(() => {
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

  return (
    <div>
      <Typography variant="h3" component="h1">
        Site Category
      </Typography>
      <Button onClick={onClickAddButtonHandler}>Add</Button>
      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <PaginationHeader
            totalCount={siteCategories.totalCount}
            page={page}
            pageSize={siteCategories.pageSize}
            dataSize={siteCategories.data.length}
          />
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Label</TableCell>
                  <TableCell>Site Count</TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {siteCategories.data.map((siteCategory) => {
                  return (
                    <TableRow key={siteCategory.id}>
                      <TableCell>{siteCategory.id}</TableCell>
                      <TableCell>{siteCategory.label}</TableCell>
                      <TableCell>{siteCategory.sites_count}</TableCell>
                      <TableCell>
                        <Button
                          onClick={onClickEditButtonHandler.bind(
                            this,
                            siteCategory
                          )}
                        >
                          Edit
                        </Button>
                      </TableCell>
                      <TableCell>
                        <IconButton
                          onClick={onClickDeleteHandler.bind(
                            this,
                            siteCategory.id
                          )}
                          aria-label="remove"
                        >
                          <DeleteIcon />
                        </IconButton>
                      </TableCell>
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          </TableContainer>

          <Pagination
            page={page}
            count={siteCategories.totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}

      {addOpen && (
        <AddSiteCategoryFormDialog
          open={addOpen}
          handleClose={handleAddClose}
        />
      )}
      {editTarget && (
        <EditSiteCategoryFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          siteCategory={editTarget}
        />
      )}
    </div>
  );
};

SiteCategories.getLayout = (page: ReactElement) => (
  <AdminLayout>{page}</AdminLayout>
);

export default SiteCategories;
