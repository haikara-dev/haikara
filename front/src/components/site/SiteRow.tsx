import React from "react";
import { Site } from "@/pages/site";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

import DeleteIcon from "@mui/icons-material/Delete";
export type SiteRowProps = {
  site: Site;
  activeSite: (id: number) => void;
  deActiveSite: (id: number) => void;
  removeSite: (id: number) => void;

  openDialog: (site: Site) => void;
  runCrawling: (id: number) => void;
  dryRunCrawling: (id: number) => void;
};

const SiteRow: React.FC<SiteRowProps> = ({
  site,
  activeSite,
  deActiveSite,
  removeSite,
  openDialog,
  runCrawling,
  dryRunCrawling,
}) => {
  const onClickRunHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    runCrawling(site.id);
  };

  const onClickDryRunHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    dryRunCrawling(site.id);
  };

  const onChangeCheckboxHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (site.active) {
      deActiveSite(site.id);
    } else {
      activeSite(site.id);
    }
  };

  const onClickRemoveHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    removeSite(site.id);
  };

  const onClickTextHandler = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault();
    openDialog(site);
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
            {site.name} （ {site.url} ）
          </Typography>
        ) : (
          <Typography variant="body1" sx={{ textDecoration: "line-through" }}>
            {site.name} （ {site.url} ）
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
