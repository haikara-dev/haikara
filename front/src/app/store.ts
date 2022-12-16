import {
  configureStore,
  PreloadedState,
  combineReducers,
} from "@reduxjs/toolkit";

import { setupListeners } from "@reduxjs/toolkit/query";
import authReducer from "@/features/auth/authSlice";
import { userApi } from "@/services/userApi";
import { adminApi } from "@/services/adminApi";

const rootReducer = combineReducers({
  auth: authReducer,
  [userApi.reducerPath]: userApi.reducer,
  [adminApi.reducerPath]: adminApi.reducer,
});

export function setupStore(preloadedState?: PreloadedState<RootState>) {
  const store = configureStore({
    reducer: rootReducer,
    preloadedState,
    // Adding the api middleware enables caching, invalidation, polling,
    // and other useful features of `rtk-query`.
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware({
        serializableCheck: {
          ignoredActions: ["auth/login"],
          ignoredPaths: ["auth.authUser"],
        },
      }).concat([userApi.middleware, adminApi.middleware]),
  });
  // optional, but required for refetchOnFocus/refetchOnReconnect behaviors
  // see `setupListeners` docs - takes an optional callback as the 2nd arg for customization
  setupListeners(store.dispatch);

  return store;
}

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof rootReducer>;

export type AppStore = ReturnType<typeof setupStore>;

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = AppStore["dispatch"];
