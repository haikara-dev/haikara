import React, { PropsWithChildren } from "react";
import { render } from "@testing-library/react";
import type { RenderOptions } from "@testing-library/react";

import { NextRouter } from "next/router";
import { RouterContext } from "next/dist/shared/lib/router-context";

// This type interface extends the default options for render from RTL, as well
// as allows the user to specify other things such as initialState, store.
interface ExtendedRenderOptions extends Omit<RenderOptions, "queries"> {
  router?: Partial<NextRouter>;
}

export function renderWithProviders(
  ui: React.ReactElement,
  { router, ...renderOptions }: ExtendedRenderOptions = {}
) {
  function Wrapper({ children }: PropsWithChildren<{}>): JSX.Element {
    if (router) {
      return (
        <RouterContext.Provider value={{ ...mockRouter, ...router }}>
          {children}
        </RouterContext.Provider>
      );
    }
    return <>{children}</>;
  }

  return { ...render(ui, { wrapper: Wrapper, ...renderOptions }) };
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
