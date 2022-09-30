import { User as AuthUser } from "@firebase/auth";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useSelector } from "react-redux";

export type User = {
  id: number;
  email: string;
  role: "admin" | "user";
};

export interface AuthState {
  currentUser: User | null;
  authUser: AuthUser | null;
  isAdmin: boolean;
}

export const initialState: AuthState = {
  currentUser: null,
  authUser: null,
  isAdmin: false,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login: (state, action: PayloadAction<AuthUser>) => {
      state.authUser = action.payload;

      // TODO APP DBにもユーザーを作成する
      // createUser(authUser);

      // TODO: ページ遷移処理を追加すること
    },
    logout: (state) => {
      state.authUser = null;

      // TODO: ページ遷移処理を追加すること
    },
    setAdmin: (state, action: PayloadAction<boolean>) => {
      state.isAdmin = action.payload;
    },
  },
});

/*
  Actions
 */
export const { login, logout, setAdmin } = authSlice.actions;

/*
  Selectors
 */
export const selectCurrentUser = (state: AuthState) => state.currentUser;
export const selectAuthUser = (state: AuthState) => state.authUser;
export const selectIsAdmin = (state: AuthState) => state.isAdmin;

/*
 Hooks
 */
export const useAuthSelector: TypedUseSelectorHook<AuthState> = useSelector;

export default authSlice.reducer;
