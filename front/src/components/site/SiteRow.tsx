import React, { useState } from "react";
import { Site } from "@/pages/site";
import { Box, Checkbox, IconButton, Typography } from "@mui/material";

import DeleteIcon from "@mui/icons-material/Delete";
export type SiteRowProps = {
  site: Site;
  activeSite: (id: number) => void;
  deActiveSite: (id: number) => void;
  removeSite: (id: number) => void;
  updateSite: (id: number, name: string, url: string, active: boolean) => void;
  openDialog: (site: Site) => void;
};

const SiteRow: React.FC<SiteRowProps> = ({
  site,
  activeSite,
  deActiveSite,
  removeSite,
  openDialog,
}) => {
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
    <Box display="flex" alignItems="center">
      <Checkbox checked={site.active} onChange={onChangeCheckboxHandler} />
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

      <IconButton onClick={onClickRemoveHandler} aria-label="remove">
        <DeleteIcon />
      </IconButton>
    </Box>
  );
};

export default SiteRow;
