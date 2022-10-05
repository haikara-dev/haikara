import React from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

import DeleteIcon from "@mui/icons-material/Delete";
import {
  DryRunResult,
  Site,
  useActiveSiteMutation,
  useDeActiveSiteMutation,
  useDeleteSiteMutation,
  useDryRunSiteCrawlingMutation,
  useRunSiteCrawlingMutation,
} from "@/services/adminApi";
import styled from "@mui/material/styles/styled";
import { useRouter } from "next/router";

export type SiteRowProps = {
  site: Site;
  openDryDialog: (result: DryRunResult) => void;
};

const SiteUrlText = styled("span")`
  word-break: break-all;
  font-size: 0.8rem;
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
    router.push(`/sites/${site.id}/edit`);
  };

  return (
    <Box
      display="flex"
      alignItems="center"
      sx={{
        backgroundColor: site.cannot_crawl ? "#f5c4c4" : "white",
      }}
    >
      <Button onClick={onClickRunHandler}>Run</Button>
      <Button onClick={onClickDryRunHandler}>Dry</Button>
      <Box
        onClick={onClickTextHandler}
        sx={{
          flexGrow: 1,
        }}
      >
        {site.active ? (
          <Typography variant="body1">
            {site.name} <SiteUrlText>( {site.url} )</SiteUrlText>
          </Typography>
        ) : (
          <Typography variant="body1" sx={{ textDecoration: "line-through" }}>
            {site.name} <SiteUrlText>( {site.url} )</SiteUrlText>
          </Typography>
        )}
      </Box>

      <div>{site.feed_url ? "RSS" : "-"}</div>

      <Checkbox checked={site.active} onChange={onChangeCheckboxHandler} />

      <IconButton onClick={onClickRemoveHandler} aria-label="remove">
        <DeleteIcon />
      </IconButton>
    </Box>
  );
};

export default SiteRow;
