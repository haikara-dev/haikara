import { User as AuthUser } from "@firebase/auth";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

import { RootState } from "@/app/store";
import { User } from "@/features/auth/authSlice";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

// Define a service using a base URL and expected endpoints
export const userApi = createApi({
  reducerPath: "userApi",
  baseQuery: fetchBaseQuery({
    baseUrl: BACKEND_API_URL,
    prepareHeaders: async (headers, { getState }) => {
      const authUser = (getState() as RootState).auth.authUser;

      if (authUser) {
        const idToken = await authUser.getIdToken();
        if (idToken) {
          headers.set("Authorization", `Bearer ${idToken}`);
        }
      }
      headers.set("Content-Type", "application/json");
      return headers;
    },
  }),
  tagTypes: ["CurrentUser"],
  endpoints: (builder) => ({
    getCurrentUser: builder.query<User, void>({
      query: () => ({
        url: `/users/current`,
      }),
      providesTags: (result) => [{ type: "CurrentUser" }],
    }),
    createUser: builder.mutation<User, AuthUser>({
      query: (queryArg) => ({
        url: `/users/create`,
        method: "POST",
        body: {
          UUID: queryArg.uid,
          email: queryArg.email,
        },
      }),
      invalidatesTags: (result, error, queryArg) => [{ type: "CurrentUser" }],
    }),
  }),
});

/*
  Hooks
 */
export const {
  useGetCurrentUserQuery,
  useLazyGetCurrentUserQuery,
  useCreateUserMutation,
} = userApi;
