import React, { useEffect } from "react";
import { useAuth } from "@/services/firebase";
import { useRouter } from "next/router";
import { selectAuthUser, selectIsAdmin } from "@/features/auth/authSlice";
import { useAppSelector } from "@/app/hooks";

export type ProtectedComponentProps = {
  children: React.ReactNode;
};
const ProtectedRouterComponent: React.FC<ProtectedComponentProps> = ({
  children,
}) => {
  // firebase auth
  const { initializedAuth } = useAuth();
  const router = useRouter();
  const authUser = useAppSelector(selectAuthUser);
  const isAdmin = useAppSelector(selectIsAdmin);

  useEffect(() => {
    if (initializedAuth) {
      console.log("authUser", authUser);
      console.log("isAdmin", isAdmin);
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
        }
      }

      if (adminRoutes.includes(router.pathname)) {
        if (!isAdmin) {
          router.push("/dashboard");
        }
      }

      // ログインしていなくても表示するページ
      // /login /register

      // ログインしていたら、トップページにリダイレクトするページ
      // /login /register

      if ("/login" === router.pathname || "/register" === router.pathname) {
        if (authUser) {
          router.push("/");
        }
      }
    }
  }, [initializedAuth]);

  return <>{initializedAuth && children}</>;
};
export default ProtectedRouterComponent;
