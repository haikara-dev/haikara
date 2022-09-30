import React from "react";
import { useAuth } from "@/services/firebase";
import { useRouter } from "next/router";
import {
  selectAuthUser,
  selectIsAdmin,
  useAuthSelector,
} from "@/features/auth/authSlice";

export type ProtectedComponentProps = {
  children: React.ReactNode;
};
const ProtectedRouterComponent: React.FC<ProtectedComponentProps> = ({
  children,
}) => {
  // firebase auth
  const { initializedAuth } = useAuth();
  const router = useRouter();
  const authUser = useAuthSelector(selectAuthUser);
  const isAdmin = useAuthSelector(selectIsAdmin);

  if (initializedAuth) {
    const adminRoutes: string[] = [
      "/sites",
      "/sites/add",
      "/sites/[id]",
      "/sites/[id]/edit",
      "/feed",
      "/user",
      "/setting",
      "/article",
    ];

    console.log("router.pathname", router.pathname);

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

  return <>{initializedAuth && children}</>;
};
export default ProtectedRouterComponent;
