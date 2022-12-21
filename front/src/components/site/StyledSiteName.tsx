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
            <span key={index} className="inline-block pr-1 font-bold">
              {name}
            </span>
          );
        }
        return (
          <span key={index} className="text-gray-300">
            /{name}
          </span>
        );
      })}
    </>
  );
};

export default StyledSiteName;
