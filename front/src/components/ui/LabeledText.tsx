import TextField from "@mui/material/TextField";
import React from "react";
import Stack from "@mui/material/Stack";

export type LabeledTextProps = {
  label: string;
  value: string;
};
const LabeledText: React.FC<LabeledTextProps> = ({ label, value }) => {
  return (
    <Stack>
      <span>{label}</span>
      <TextField value={value} aria-readonly={true} />
    </Stack>
  );
};

export default LabeledText;
