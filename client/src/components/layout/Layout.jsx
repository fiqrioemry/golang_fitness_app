import Header from "./Header";
import Footer from "./Footer";
import { Fragment } from "react";
import { Outlet } from "react-router-dom";

const Layout = () => {
  return (
    <Fragment>
      <Header />
      <Outlet />
      <Footer />
    </Fragment>
  );
};

export default Layout;
