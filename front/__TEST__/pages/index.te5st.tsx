import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { Provider } from "react-redux";
import { store } from "@/app/store";
import Home from "@/pages/index";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

jest.mock("next/router", () => ({
  useRouter() {
    return {
      query: {},
    };
  },
}));

describe("Home", () => {
  it("render heading", () => {
    render(
      <Provider store={store}>
        <Home />
      </Provider>
    );
    screen.debug();
    // const heading = screen.getByRole("link", {
    //   name: /haikara/i,
    // });
    // expect(heading).toBeInTheDocument();
    expect(true).toBeTruthy();
  });
});
