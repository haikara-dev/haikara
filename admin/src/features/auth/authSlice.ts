import { User as AuthUser } from "@firebase/auth";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "@/app/store";

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
    },
    logout: () => initialState,
    setCurrentUser: (state, action: PayloadAction<User>) => {
      state.currentUser = action.payload;
      state.isAdmin = action.payload.role === "admin";
    },
    setAdmin: (state, action: PayloadAction<boolean>) => {
      state.isAdmin = action.payload;
    },
  },
});

/*
  Actions
 */
export const { login, logout, setCurrentUser, setAdmin } = authSlice.actions;

/*
  Selectors
 */

export const selectCurrentUser = (state: RootState) => state.auth.currentUser;
export const selectAuthUser = (state: RootState) => state.auth.authUser;
export const selectIsAdmin = (state: RootState) => state.auth.isAdmin;

export default authSlice.reducer;
