import React from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";

import DeleteIcon from "@mui/icons-material/Delete";
import {
  Site,
  useActiveSiteMutation,
  useDeActiveSiteMutation,
  useDeleteSiteMutation,
  useDryRunSiteCrawlingMutation,
  useLazyGetSiteWithSiteCrawlRuleQuery,
  useRunSiteCrawlingMutation,
} from "@/services/adminApi";

export type SiteRowProps = {
  site: Site;
  openDialog: (site: Site) => void;
};

const SiteRow: React.FC<SiteRowProps> = ({ site, openDialog }) => {
  const [runSiteCrawling] = useRunSiteCrawlingMutation();
  const [dryRunSiteCrawling] = useDryRunSiteCrawlingMutation();
  const [activeSite] = useActiveSiteMutation();
  const [deActiveSite] = useDeActiveSiteMutation();
  const [deleteSite] = useDeleteSiteMutation();

  const [getSiteWithSiteCrawlRule] = useLazyGetSiteWithSiteCrawlRuleQuery();

  const onClickGetHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    getSiteWithSiteCrawlRule(site.id);
  };

  const onClickRunHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    runSiteCrawling(site.id);
  };

  const onClickDryRunHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    dryRunSiteCrawling(site.id);
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
      <Button onClick={onClickGetHandler}>Get</Button>
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
