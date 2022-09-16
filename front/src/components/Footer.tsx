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
          href="https://haikara-dev.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          haikara-dev
        </a>
      </Container>
    </Box>
  );
};

export default Footer;
