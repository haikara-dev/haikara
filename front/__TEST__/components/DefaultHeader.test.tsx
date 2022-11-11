import { render, screen } from "@testing-library/react";
import DefaultHeader from "@/components/DefaultHeader";
import "@testing-library/jest-dom";
import { Provider } from "react-redux";
import { store } from "@/app/store";

jest.mock("firebase/auth", () => {
  return {
    getAuth: jest.fn(),
  };
});

describe("DefaultHeader", () => {
  it("render heading", () => {
    render(
      <Provider store={store}>
        <DefaultHeader />
      </Provider>
    );
    const heading = screen.getByRole("link", {
      name: /haikara/i,
    });
    expect(heading).toBeInTheDocument();
  });
});
