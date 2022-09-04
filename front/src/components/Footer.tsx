import { Box, Container } from "@mui/material";

const Footer = () => {
  return (
    <Box component="footer">
      <Container
        sx={{
          textAlign: "center",
        }}
      >
        Powered by{" "}
        <a
          href="https://cubdesign.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          cubdesign
        </a>
      </Container>
    </Box>
  );
};

export default Footer;
