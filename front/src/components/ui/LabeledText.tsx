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
      <span role={"label"}>{label}</span>
      <TextField
        value={value}
        aria-readonly={true}
        sx={{
          backgroundColor: "#ddf1f6",
        }}
      />
    </Stack>
  );
};

export default LabeledText;
