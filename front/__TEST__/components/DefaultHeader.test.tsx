import { screen } from "@testing-library/react";
import DefaultHeader from "@/components/DefaultHeader";
import "@testing-library/jest-dom";
import { renderWithProviders } from "../utils/test-utils";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

// jest.mock("@/app/hooks", () => {
//   return {
//     useAppDispatch: jest.fn(),
//     useAppSelector: jest.fn().mockImplementation((selecter) => selecter()),
//   };
// });
//
// jest.mock("@/features/auth/authSlice", () => {
//   return {
//     logout: jest.fn(),
//     selectAuthUser: jest.fn().mockReturnValue(true),
//   };
// });

describe("DefaultHeader", () => {
  it("render heading", () => {
    renderWithProviders(<DefaultHeader />);

    const heading = screen.getByRole("link", {
      name: /haikara/i,
    });

    expect(heading).toBeInTheDocument();

    const login = screen.getByText(/ログイン/i);
    expect(login).toBeInTheDocument();
  });

  it("render heading with Login", () => {
    renderWithProviders(<DefaultHeader />, {
      preloadedState: {
        auth: {
          currentUser: {
            id: 1,
            email: "a@b",
            role: "user",
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
          isAdmin: false,
        },
      },
    });

    const link = screen.getByText(/コンソール/i);
    expect(link).toBeInTheDocument();

    const logout = screen.getByText(/ログアウト/i);
    expect(logout).toBeInTheDocument();
  });
});
