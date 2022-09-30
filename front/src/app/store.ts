import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import counterReducer from "@/features/counter/counterSlice";
import authReducer from "@/features/auth/authSlice";
import { pokemonApi } from "@/services/pokemon";
import { userApi } from "@/services/userApi";
export const store = configureStore({
  reducer: {
    counter: counterReducer,
    auth: authReducer,
    [userApi.reducerPath]: userApi.reducer,
    [pokemonApi.reducerPath]: pokemonApi.reducer,
  },
  // Adding the api middleware enables caching, invalidation, polling,
  // and other useful features of `rtk-query`.
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat([userApi.middleware, pokemonApi.middleware]),
});
// optional, but required for refetchOnFocus/refetchOnReconnect behaviors
// see `setupListeners` docs - takes an optional callback as the 2nd arg for customization
setupListeners(store.dispatch);

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
