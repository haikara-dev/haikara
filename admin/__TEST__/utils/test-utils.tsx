import React, { PropsWithChildren } from "react";
import { render } from "@testing-library/react";
import type { RenderOptions } from "@testing-library/react";
import type { PreloadedState } from "@reduxjs/toolkit";
import { Provider } from "react-redux";

import type { AppStore, RootState } from "@/app/store";
import { setupStore } from "@/app/store";
import { NextRouter } from "next/router";
import { RouterContext } from "next/dist/shared/lib/router-context";

// This type interface extends the default options for render from RTL, as well
// as allows the user to specify other things such as initialState, store.
interface ExtendedRenderOptions extends Omit<RenderOptions, "queries"> {
  router?: Partial<NextRouter>;
  preloadedState?: PreloadedState<RootState>;
  store?: AppStore;
}

export function renderWithProviders(
  ui: React.ReactElement,
  {
    router,
    preloadedState = {},
    // Automatically create a store instance if no store was passed in
    store = setupStore(preloadedState),
    ...renderOptions
  }: ExtendedRenderOptions = {}
) {
  function Wrapper({ children }: PropsWithChildren<{}>): JSX.Element {
    if (router) {
      return (
        <RouterContext.Provider value={{ ...mockRouter, ...router }}>
          <Provider store={store}>{children}</Provider>{" "}
        </RouterContext.Provider>
      );
    }
    return <Provider store={store}>{children}</Provider>;
  }

  return { store, ...render(ui, { wrapper: Wrapper, ...renderOptions }) };
}

const mockRouter: NextRouter = {
  basePath: "",
  route: "/",
  pathname: "/",
  query: {},
  asPath: "/",
  isLocaleDomain: true,
  isReady: true,
  push: jest.fn(),
  prefetch: jest.fn(),
  replace: jest.fn(),
  reload: jest.fn(),
  back: jest.fn(),
  beforePopState: jest.fn(),
  events: {
    on: jest.fn(),
    off: jest.fn(),
    emit: jest.fn(),
  },
  isFallback: false,
  isPreview: false,
};

export function mockAuthorizedAuth(role: "user" | "admin" = "user") {
  return {
    auth: {
      currentUser: {
        id: 1,
        email: "a@b",
        role: role,
      },
      authUser: {
        displayName: "a",
        email: "a@b",
        phoneNumber: null,
        photoURL: null,
        providerId: "",
        uid: "1",

        emailVerified: false,
        isAnonymous: false,
        metadata: {},
        providerData: [],
        refreshToken: "",
        tenantId: null,

        delete: jest.fn(),
        getIdToken: jest.fn(),
        getIdTokenResult: jest.fn(),
        reload: jest.fn(),
        toJSON: jest.fn(),
      },
      isAdmin: role === "admin",
    },
  };
}
