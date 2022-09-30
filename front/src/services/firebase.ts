import { useEffect, useState } from "react";
import { useAppDispatch } from "@/app/hooks";
import { Auth, getAuth } from "firebase/auth";

export const useAuth = () => {
  const auth: Auth = getAuth();
  const dispatch = useAppDispatch();
  const [initializedAuth, setInitializedAuth] = useState<boolean>(false);
  useEffect(() => {
    const unSubscribe = auth.onAuthStateChanged((authUser) => {
      // setAuthUser(authUser);
      if (authUser) {
        // getUser(authUser);
      } else {
        // setInitialize(true);
      }
      // TODO ここでいいのか？
      setInitializedAuth(true);
    });
    return () => {
      // cleanup
      unSubscribe();
    };
  }, []);

  return { initializedAuth };
};
