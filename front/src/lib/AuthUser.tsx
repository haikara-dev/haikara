import { Auth, getAuth, User as AuthUser } from "firebase/auth";
import React, { createContext, useContext, useEffect, useState } from "react";
import { useRouter } from "next/router";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export type User = {
  id: number;
  email: string;
  role: "admin" | "user";
};

export type AuthUserContextType = {
  currentUser: User | null;
  authUser: AuthUser | null;
  login: (user: AuthUser, callback: () => void) => void;
  logout: (callback: () => void) => void;
  isAdmin: boolean;
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
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [initialize, setInitialize] = useState<boolean>(false);
  const [isAdmin, setIsAdmin] = useState<boolean>(false);

  const router = useRouter();

  useEffect(() => {
    const unSubscribe = auth.onAuthStateChanged((authUser) => {
      console.log("onAuthStateChanged", authUser);
      setAuthUser(authUser);
      if (authUser) {
        getUser(authUser);
      } else {
        setInitialize(true);
      }
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
    currentUser,
    authUser,
    login,
    logout,
    isAdmin,
  };

  const getRequestHeaders = async (authUser: AuthUser) => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const createUser = async (authUser: AuthUser) => {
    try {
      const headers = await getRequestHeaders(authUser);
      const uuid = authUser.uid;
      const email = authUser.email;
      const res = await fetch(BACKEND_API_URL + "/users/create", {
        method: "POST",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify({
          UUID: uuid,
          email: email,
        }),
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
    } catch (err) {
      console.log(err);
    }
  };

  const getUser = async (authUser: AuthUser) => {
    try {
      const headers = await getRequestHeaders(authUser);
      const res = await fetch(BACKEND_API_URL + "/users/current", {
        method: "GET",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setCurrentUser({
        id: json.id,
        email: json.email,
        role: json.role,
      });
      setIsAdmin(json.role === "admin");
      setInitialize(true);
    } catch (err) {
      console.log(err);
    }
  };

  if (initialize) {
    const adminRoutes: string[] = [
      "/site",
      "/feed",
      "/user",
      "/setting",
      "/article",
    ];
    const authorizedRoutes: string[] = [...adminRoutes, "/dashboard"];

    if (authorizedRoutes.includes(router.pathname)) {
      if (!authUser) {
        router.push("/login");
        return <></>;
      }
    }

    if (adminRoutes.includes(router.pathname)) {
      if (!isAdmin) {
        router.push("/dashboard");
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
