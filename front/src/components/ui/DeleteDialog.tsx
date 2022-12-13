import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import DialogContent from "@mui/material/DialogContent";
import { FC, ReactNode } from "react";

export type DeleteDialogProps = {
  open: boolean;
  title: string;
  children: ReactNode;

  cancelButtonLabel?: string;
  agreeButtonLabel?: string;

  cancelHandler: () => void;
  agreeHandler: () => void;
};

const DeleteDialog: FC<DeleteDialogProps> = ({
  open,
  title,
  children,
  cancelButtonLabel = "Chancel",
  agreeButtonLabel = "Agree",
  cancelHandler,
  agreeHandler,
}) => {
  return (
    <Dialog
      open={open}
      onClose={cancelHandler}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">{title}</DialogTitle>
      <DialogContent>
        <DialogContentText id="alert-dialog-description">
          {children}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={cancelHandler} data-testid="cancel-button">
          {cancelButtonLabel}
        </Button>
        <Button onClick={agreeHandler} data-testid="agree-button" autoFocus>
          {agreeButtonLabel}
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default DeleteDialog;
