// src/App.jsx
import Home from "./pages/Home";
import NotFound from "./pages/NotFound";
import Profile from "./pages/customer/Profile";
import CreateClass from "./pages/classes/CreateClass";
import CreatePackage from "./pages/packages/CreatePackage";
import OptionsDisplay from "./pages/options/OptionsDisplay";
import ClassesDisplay from "./pages/classes/ClassesDisplay";
import PackagesDisplay from "./pages/packages/PackagesDisplay";

import { useEffect } from "react";
import { Toaster } from "sonner";
import Layout from "./components/layout/Layout";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "./store/useAuthStore";
import { AuthRoute, NonAuthRoute } from "./middleware";
import UserLayout from "./components/layout/UserLayout";
import { Navigate, Route, Routes } from "react-router-dom";

function App() {
  const { checkingAuth, authMe } = useAuthStore();

  useEffect(() => {
    authMe();
  }, []);

  if (checkingAuth) return <Loading />;

  return (
    <>
      <Toaster />
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route
            path="user"
            element={
              <AuthRoute>
                <UserLayout />
              </AuthRoute>
            }
          >
            <Route path="profile" element={<Profile />} />
            <Route index element={<Navigate to="profile" replace />} />
          </Route>
        </Route>

        <Route path="/options" element={<OptionsDisplay />} />
        <Route path="/classes" element={<ClassesDisplay />} />
        <Route path="/classes/add" element={<CreateClass />} />
        <Route path="/packages" element={<PackagesDisplay />} />
        <Route path="/packages/add" element={<CreatePackage />} />

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
