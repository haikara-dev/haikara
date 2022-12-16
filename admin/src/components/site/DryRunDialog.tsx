import React from "react";
import Highlight, { defaultProps } from "prism-react-renderer";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import Dialog from "@mui/material/Dialog";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import { DryRunResult } from "@/services/adminApi";

export type AddSiteFormProps = {
  open: boolean;
  handleClose: () => void;
  dryRunResult: DryRunResult;
};

const DryRunDialog: React.FC<AddSiteFormProps> = ({
  open,
  handleClose,
  dryRunResult,
}) => {
  return (
    <Dialog open={open} onClose={handleClose} maxWidth="xl" fullWidth={true}>
      <DialogTitle>Dry Run</DialogTitle>
      <DialogContent>
        <DialogContentText>結果を表示します。</DialogContentText>
        <Box display="flex" flexDirection="column" gap={2}>
          {dryRunResult.count}
          <Highlight
            {...defaultProps}
            code={dryRunResult.contents}
            language="jsx"
          >
            {({ className, style, tokens, getLineProps, getTokenProps }) => (
              <pre className={className} style={style}>
                {tokens.map((line, i) => (
                  <div key={i} {...getLineProps({ line })}>
                    {line.map((token, key) => (
                      <span key={key} {...getTokenProps({ token })} />
                    ))}
                  </div>
                ))}
              </pre>
            )}
          </Highlight>
        </Box>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Close</Button>
      </DialogActions>
    </Dialog>
  );
};

export default DryRunDialog;
