import { Auth, getAuth, User as AuthUser } from "firebase/auth";
import React, { createContext, useContext, useEffect, useState } from "react";
import { useRouter } from "next/router";

export type AuthUserContextType = {
  authUser: AuthUser | null;
  login: (user: AuthUser, callback: () => void) => void;
  logout: (callback: () => void) => void;
};

const AuthUserContext = createContext<AuthUserContextType>(
  {} as AuthUserContextType
);

export const useAuthUserContext = (): AuthUserContextType => {
  return useContext<AuthUserContextType>(AuthUserContext);
};

type AuthUserProviderProps = {
  children: React.ReactNode;
};

const AuthUserProvider: React.FC<AuthUserProviderProps> = ({ children }) => {
  const auth: Auth = getAuth();
  const [authUser, setAuthUser] = useState<AuthUser | null>(null);
  const [initialize, setInitialize] = useState<boolean>(false);

  const router = useRouter();
  console.log(router.pathname);

  useEffect(() => {
    const unSubscribe = auth.onAuthStateChanged((authUser) => {
      console.log("onAuthStateChanged", authUser);
      setInitialize(true);
      setAuthUser(authUser);
    });
    return () => {
      // cleanup
      unSubscribe();
    };
  }, []);

  const login = (authUser: AuthUser, callback: () => void) => {
    setAuthUser(authUser);
    createUser(authUser);
    callback();
  };

  const logout = (callback: () => void) => {
    setAuthUser(null);
    callback();
  };

  const value: AuthUserContextType = {
    authUser,
    login,
    logout,
  };

  const getRequestHeaders = async (authUser: AuthUser) => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const createUser = async (authUser: AuthUser) => {
    const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
    try {
      const headers = await getRequestHeaders(authUser);
      const uuid = authUser.uid;
      const res = await fetch(BACKEND_API_URL + "/user/create", {
        method: "POST",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify({
          UUID: uuid,
        }),
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
    } catch (err) {
      console.log(err);
    }
  };

  if (initialize) {
    if ("/todo" === router.pathname) {
      if (!authUser) {
        router.push("/login");
        return <></>;
      }
    }
    // ログインしていなくても表示するページ
    // /login /register

    // ログインしていたら、トップページにリダイレクトするページ
    // /login /register

    if ("/login" === router.pathname || "/register" === router.pathname) {
      if (authUser) {
        router.push("/");
        return <></>;
      }
    }
  }

  return (
    <AuthUserContext.Provider value={value}>
      {initialize && children}
    </AuthUserContext.Provider>
  );
};

export default AuthUserProvider;
