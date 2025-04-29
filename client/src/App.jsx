// src/App.jsx
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import NotFound from "./pages/NotFound";
import Profile from "./pages/customer/Profile";

import { Toaster } from "sonner";
import Loading from "@/components/ui/Loading";
import ErrorDialog from "@/components/ui/ErrorDialog";
import Layout from "./components/layout/Layout";
import { useAuthMe } from "./hooks/useAuthQuery";
import { AuthRoute, NonAuthRoute } from "./middleware";
import UserLayout from "./components/layout/UserLayout";
import { Navigate, Route, Routes } from "react-router-dom";
import Categories from "./pages/testing/Categories";
import Attendance from "./pages/testing/Attendance";
import Instructor from "./services/instructor";
import Class from "./services/class";
import Booking from "./pages/testing/Booking";
import Level from "./pages/testing/Level";
import { Package } from "lucide-react";
import Payment from "./services/payment";
import Review from "./services/review";
import Schedule from "./services/schedule";
import Subcategories from "./pages/testing/Subcategories";
import Type from "./services/type";

function App() {
  // const { isError, isLoading, refetch } = useAuthMe();

  // if (isLoading) return <Loading />;

  // if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <>
      <Toaster />
      <Routes>
        <Route path="/attendance" element={<Attendance />} />
        <Route path="/booking" element={<Booking />} />
        <Route path="/categories" element={<Categories />} />
        <Route path="/class" element={<Class />} />
        <Route path="/instructor" element={<Instructor />} />
        <Route path="/level" element={<Level />} />
        <Route path="/location" element={<Location />} />
        <Route path="/package" element={<Package />} />
        <Route path="/payment" element={<Payment />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/review" element={<Review />} />
        <Route path="/schedule" element={<Schedule />} />
        <Route path="/subcategories" element={<Subcategories />} />
        <Route path="/type" element={<Type />} />
        <Route
          path="/signin"
          element={
            <NonAuthRoute>
              <SignIn />
            </NonAuthRoute>
          }
        />
        <Route
          path="/signup"
          element={
            <NonAuthRoute>
              <SignUp />
            </NonAuthRoute>
          }
        />

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
