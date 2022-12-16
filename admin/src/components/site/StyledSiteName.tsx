import Box from "@mui/material/Box";

export type StyledSiteNameProps = {
  children: React.ReactNode;
};

const StyledSiteName: React.FC<StyledSiteNameProps> = ({ children }) => {
  const siteName = children!.toString();
  const siteNameArray = siteName.split("/");
  return (
    <>
      {siteNameArray.map((name, index) => {
        if (index === 0) {
          return (
            <Box
              component="span"
              key={index}
              sx={{
                display: "inline-block",
                paddingRight: 1,
                fontWeight: "bold",
              }}
            >
              {name}
            </Box>
          );
        }
        return (
          <Box
            component="span"
            key={index}
            sx={{
              color: "#999999",
            }}
          >
            /{name}
          </Box>
        );
      })}
    </>
  );
};

export default StyledSiteName;
