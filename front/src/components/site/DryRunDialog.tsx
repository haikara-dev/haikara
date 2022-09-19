import React from "react";
import Highlight, { defaultProps } from "prism-react-renderer";
import {
  Box,
  Button,
  DialogTitle,
  DialogContent,
  Dialog,
  DialogContentText,
  DialogActions,
} from "@mui/material";
import { DryRunResult } from "@/pages/site";

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
