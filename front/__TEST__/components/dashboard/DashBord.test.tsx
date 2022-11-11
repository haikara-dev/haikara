import { render, screen } from "@testing-library/react";
import DashBord from "@/components/dashboard/DashBord";
import "@testing-library/jest-dom";
import { store } from "@/app/store";
import { Provider } from "react-redux";

describe("DashBord", () => {
  it("render DashBord", () => {
    render(
      <Provider store={store}>
        <DashBord />
      </Provider>
    );
    const footer = screen.getByText(/Loading/i);
    expect(footer).toBeInTheDocument();
  });
});
