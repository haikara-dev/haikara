import { useEffect, useState } from "react";
import { useAppDispatch } from "@/app/hooks";
import { Auth, getAuth } from "firebase/auth";
import { login, logout, setCurrentUser } from "@/features/auth/authSlice";
import { userApi } from "@/services/userApi";

export const useAuth = () => {
  const auth: Auth = getAuth();
  const dispatch = useAppDispatch();
  const [initializedAuth, setInitializedAuth] = useState<boolean>(false);

  const [isLogin, setIsLogin] = useState<boolean>(false);

  useEffect(() => {
    const unSubscribe = auth.onAuthStateChanged(async (_authUser) => {
      if (_authUser) {
        dispatch(login(_authUser));
        setIsLogin(true);
      } else {
        dispatch(logout());
        setInitializedAuth(true);
      }
    });

    return () => {
      // cleanup
      unSubscribe();
    };
  }, []);

  useEffect(() => {
    const f = async () => {
      const { data: currentUser, isSuccess } = await dispatch(
        userApi.endpoints.getCurrentUser.initiate()
      );

      if (isSuccess) {
        dispatch(setCurrentUser(currentUser));
      }

      setInitializedAuth(true);
    };
    if (isLogin) {
      f();
    }
  }, [isLogin]);
  return { initializedAuth };
};