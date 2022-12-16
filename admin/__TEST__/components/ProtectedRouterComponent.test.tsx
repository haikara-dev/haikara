import { screen } from "@testing-library/react";
import ProtectedRouterComponent from "@/components/ProtectedRouterComponent";
import { mockAuthorizedAuth, renderWithProviders } from "../utils/test-utils";

jest.mock("@/services/firebase", () => ({
  useAuth: () => {
    return {
      initializedAuth: true,
    };
  },
}));

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("ProtectedRouterComponent", () => {
  it("renders", () => {
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/",
          push: jest.fn(),
        },
      }
    );

    expect(screen.getByText("Test")).toBeInTheDocument();
  });
});

describe("ProtectedRouterComponent guest", () => {
  it("renders", () => {
    const pushFn = jest.fn();
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/sites",
          push: pushFn,
        },
        preloadedState: {
          auth: {
            currentUser: null,
            authUser: null,
            isAdmin: false,
          },
        },
      }
    );

    expect(pushFn).toHaveBeenCalledWith("/login");
  });
});

describe("ProtectedRouterComponent authorized", () => {
  it("/login ログインしている場合、/ に遷移する", () => {
    const pushFn = jest.fn();
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/login",
          push: pushFn,
        },
        preloadedState: { ...mockAuthorizedAuth("user") },
      }
    );

    expect(pushFn).toHaveBeenCalledWith("/");
  });

  it("/register ログインしている場合、/ に遷移する", () => {
    const pushFn = jest.fn();
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/register",
          push: pushFn,
        },
        preloadedState: { ...mockAuthorizedAuth("user") },
      }
    );

    expect(pushFn).toHaveBeenCalledWith("/");
  });
});

describe("ProtectedRouterComponent authorized user", () => {
  it("アクセスできないパスの場合、/dashboard に遷移する", () => {
    const pushFn = jest.fn();
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/sites",
          push: pushFn,
        },
        preloadedState: { ...mockAuthorizedAuth("user") },
      }
    );

    expect(pushFn).toHaveBeenCalledWith("/dashboard");
  });
});

describe("ProtectedRouterComponent authorized admin", () => {
  it("何もしない", () => {
    const pushFn = jest.fn();
    renderWithProviders(
      <ProtectedRouterComponent>Test</ProtectedRouterComponent>,
      {
        router: {
          pathname: "/sites",
          push: pushFn,
        },
        preloadedState: { ...mockAuthorizedAuth("admin") },
      }
    );

    expect(pushFn).not.toHaveBeenCalled();
  });
});
