import Box from "@mui/material/Box";
import Container from "@mui/material/Container";

const Footer = () => {
  return (
    <Box component="footer">
      <Container
        sx={{
          textAlign: "center",
        }}
      >
        <a href="https://haikara.dev" target="_blank" rel="noopener noreferrer">
          ©︎haikara
        </a>
      </Container>
    </Box>
  );
};

export default Footer;
