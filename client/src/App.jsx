// src/App.jsx
import Home from "./pages/Home";
// import SignIn from "./pages/SignIn";
// import SignUp from "./pages/SignUp";
import NotFound from "./pages/NotFound";
import Profile from "./pages/customer/Profile";

import { Toaster } from "sonner";
import Layout from "./components/layout/Layout";
import { useAuthMe } from "./hooks/useAuthQuery";
import { Loading } from "@/components/ui/Loading";
import { AuthRoute, NonAuthRoute } from "./middleware";
import UserLayout from "./components/layout/UserLayout";
import { Navigate, Route, Routes } from "react-router-dom";
import ClassDisplay from "./pages/testing/ClassDisplay";

function App() {
  const { isLoading } = useAuthMe();

  if (isLoading) return <Loading />;

  return (
    <>
      <Toaster />
      <Routes>
        <Route path="/class" element={<ClassDisplay />} />
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

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
