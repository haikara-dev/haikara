import React from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";

import DeleteIcon from "@mui/icons-material/Delete";
import {
  DryRunResult,
  Site,
  SiteWithCategory,
  useActiveSiteMutation,
  useDeActiveSiteMutation,
  useDeleteSiteMutation,
  useDryRunSiteCrawlingMutation,
  useRunSiteCrawlingMutation,
} from "@/services/adminApi";
import styled from "@mui/material/styles/styled";
import { useRouter } from "next/router";
import StyledSiteName from "@/components/site/StyledSiteName";

export type SiteRowProps = {
  site: SiteWithCategory;
  openDryDialog: (result: DryRunResult) => void;
};

const SiteUrlText = styled("span")`
  display: block;
  word-break: break-all;
  font-size: 0.8rem;
  color: #999999;
`;

const SiteRow: React.FC<SiteRowProps> = ({ site, openDryDialog }) => {
  const [runSiteCrawling] = useRunSiteCrawlingMutation();
  const [dryRunSiteCrawling] = useDryRunSiteCrawlingMutation();
  const [activeSite] = useActiveSiteMutation();
  const [deActiveSite] = useDeActiveSiteMutation();
  const [deleteSite] = useDeleteSiteMutation();

  const router = useRouter();

  const onClickRunHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    runSiteCrawling(site.id);
  };

  const onClickDryRunHandler = async (
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    const result = await dryRunSiteCrawling(site.id).unwrap();
    openDryDialog(result);
  };

  const onChangeCheckboxHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (site.active) {
      deActiveSite({
        id: site.id,
        body: {
          active: false,
        },
      });
    } else {
      activeSite({
        id: site.id,
        body: {
          active: true,
        },
      });
    }
  };

  const onClickRemoveHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    deleteSite(site.id);
  };

  const onClickTextHandler = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault();
    router.push(`/sites/${site.id}`);
  };

  const onClickEditHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    router.push(`/sites/${site.id}/edit`);
  };

  return (
    <TableRow
      sx={{
        backgroundColor: site.cannot_crawl ? "#f5c4c4" : "white",
      }}
    >
      <TableCell>
        <Button onClick={onClickRunHandler}>Run</Button>
      </TableCell>
      <TableCell>
        <Button onClick={onClickDryRunHandler}>Dry</Button>
      </TableCell>
      <TableCell>
        <Box
          onClick={onClickTextHandler}
          sx={{
            flexGrow: 1,
          }}
        >
          {site.active ? (
            <Typography variant="body1">
              <StyledSiteName>{site.name}</StyledSiteName>
              <SiteUrlText>{site.url}</SiteUrlText>
            </Typography>
          ) : (
            <Typography variant="body1" sx={{ textDecoration: "line-through" }}>
              <StyledSiteName>{site.name}</StyledSiteName>
              <SiteUrlText>{site.url}</SiteUrlText>
            </Typography>
          )}
        </Box>
      </TableCell>
      <TableCell>
        {site.site_categories.map((category) => (
          <Typography
            key={category.id}
            variant="body1"
            sx={{
              marginRight: 1,
              whiteSpace: "nowrap",
            }}
          >
            {category.label}
          </Typography>
        ))}
      </TableCell>
      <TableCell>{site.feed_url ? "RSS" : "-"}</TableCell>
      <TableCell>
        <Button onClick={onClickEditHandler}>Edit</Button>
      </TableCell>
      <TableCell>
        <Checkbox checked={site.active} onChange={onChangeCheckboxHandler} />
      </TableCell>
      <TableCell>
        <IconButton onClick={onClickRemoveHandler} aria-label="remove">
          <DeleteIcon />
        </IconButton>
      </TableCell>
    </TableRow>
  );
};

export default SiteRow;
