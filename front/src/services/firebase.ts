import { useEffect, useState } from "react";
import { useAppDispatch, useAppSelector } from "@/app/hooks";
import { Auth, getAuth } from "firebase/auth";
import {
  login,
  logout,
  selectAuthUser,
  setCurrentUser,
} from "@/features/auth/authSlice";
import { userApi } from "@/services/userApi";

export const useAuth = () => {
  const auth: Auth = getAuth();
  const dispatch = useAppDispatch();
  const [initializedAuth, setInitializedAuth] = useState<boolean>(false);

  const [isLogin, setIsLogin] = useState<boolean>(false);

  const authUser = useAppSelector(selectAuthUser);

  useEffect(() => {
    const unSubscribe = auth.onAuthStateChanged(async (_authUser) => {
      console.log("xxxxxxxxxxx", _authUser);

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
        userApi.endpoints.getCurrentUser.initiate(authUser!)
      );

      if (isSuccess) {
        dispatch(setCurrentUser(currentUser));
      }

      console.log("dddddddddddddddddddddddd 2", isLogin);
      setInitializedAuth(true);
    };
    if (isLogin) {
      console.log("dddddddddddddddddddddddd 1", isLogin);
      f();
    }
  }, [isLogin]);
  return { initializedAuth };
};
