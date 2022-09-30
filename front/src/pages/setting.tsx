import React, { ReactElement, useCallback, useState } from "react";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import Highlight, { defaultProps } from "prism-react-renderer";
import { useDropzone } from "react-dropzone";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";
import { selectAuthUser, useAuthSelector } from "@/features/auth/authSlice";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const Settings: NextPageWithLayout = () => {
  const [logs, setLogs] = useState<string[]>([]);
  const [importFile, setImportFile] = useState<File | null>(null);

  const authUser = useAuthSelector(selectAuthUser);

  const onDrop = useCallback((acceptedFiles: any) => {
    // Do something with the files
    setImportFile(acceptedFiles[0]);
  }, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };
  const exportSite = async () => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites/export", {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      // TO "FORCE DOWNLOAD"
      const url = window.URL.createObjectURL(await res.blob());
      const a = document.createElement("a");
      a.href = url;
      a.download = "export-sites.json";
      a.click();

      // CLEAN UP
      window.URL.revokeObjectURL(url);
      document.removeChild(a);

      // const json = await res.json();
      // setLogs([
      //   ...logs,
      //   `${new Date().toLocaleString()} :\n\n ${JSON.stringify(
      //     json,
      //     null,
      //     "\t"
      //   )}`,
      // ]);
    } catch (err) {
      console.log(err);
    }
  };
  const importSite = async () => {
    try {
      if (importFile == null) {
        return;
      }

      const formData = new FormData();
      formData.append("file", importFile);

      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites/import", {
        method: "POST",
        headers: headers,
        body: formData,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setLogs([
        ...logs,
        `${new Date().toLocaleString()} :\n\n ${JSON.stringify(
          json,
          null,
          "\t"
        )}`,
      ]);
    } catch (err) {
      console.log(err);
    }
  };
  const handleExportButtonClick = () => {
    exportSite();
  };

  const handleImportButtonClick = () => {
    importSite();
  };

  return (
    <div>
      <Typography variant="h3" component="h1">
        Settings
      </Typography>

      <h2>Site</h2>

      <Stack spacing={2}>
        <Card>
          <Stack direction="row" alignItems="center">
            <div>サイトをjson形式でダウンロードします。</div>
            <Button onClick={handleExportButtonClick}>エクスポート</Button>
          </Stack>
        </Card>

        <Card>
          <Stack direction="row" alignItems="center">
            <div>サイトをjson形式でアップロードし、インポートします。</div>

            <Button onClick={handleImportButtonClick}>インポート</Button>
          </Stack>
          <div {...getRootProps()}>
            <input {...getInputProps()} />
            {importFile ? (
              <Box
                p={3}
                sx={{
                  backgroundColor: "grey.100",
                }}
              >
                {importFile.name}
              </Box>
            ) : isDragActive ? (
              <Box
                p={3}
                sx={{
                  backgroundColor: "grey.100",
                }}
              >
                Drop the files here ...
              </Box>
            ) : (
              <Box
                p={3}
                sx={{
                  backgroundColor: "grey.100",
                }}
              >
                Drag &apos;n&apos; drop some files here, or click to select
                files
              </Box>
            )}
          </div>{" "}
        </Card>
      </Stack>

      <div>
        {logs.map((log, index) => {
          return (
            <Highlight {...defaultProps} key={index} code={log} language="json">
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
          );
        })}
      </div>
    </div>
  );
};

Settings.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Settings;
