// public pages
import Home from "./pages/Home";
import About from "./pages/About";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import Classes from "./pages/Classes";
import NotFound from "./pages/NotFound";
import Schedules from "./pages/Schedules";
import ClassDetail from "./pages/ClassDetail";
import Profile from "./pages/customer/Profile";

// admin pages

import Dashboard from "./pages/admin/Dashboard";
import UsersList from "./pages/admin/UsersList";
import AddClass from "./pages/admin/classes/AddClass";
import BookingsList from "./pages/admin/BookingsList";
import VouchersList from "./pages/admin/VouchersList";
import ReviewsLists from "./pages/admin/ReviewsLists";
import AddPackage from "./pages/admin/packages/AddPackage";
import ClassesList from "./pages/admin/classes/ClassesList";
import ClassOptions from "./pages/admin/classes/ClassOptions";
import TransactionsList from "./pages/admin/TransactionsList";
import PackagesList from "./pages/admin/packages/PackagesList";
import NotificationsList from "./pages/admin/NotificationsList";
import ClassSchedules from "./pages/admin/classes/ClassSchedules";
import AddInstructors from "./pages/admin/instructors/AddInstructors";
import InstructorsList from "./pages/admin/instructors/InstructorsList";

// customer pages
import Packages from "./pages/Packages";
import PackageDetail from "./pages/PackageDetail";
import UserBookings from "./pages/customer/UserBookings";
import UserPackages from "./pages/customer/UserPackages";
import UserTransactions from "./pages/customer/UserTransactions";
import UserNotifications from "./pages/customer/UserNotifications";

// route config & support
import { Toaster } from "sonner";
import { useEffect } from "react";
import ScrollToTop from "./hooks/useScrollToTop";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "./store/useAuthStore";
import { Navigate, Route, Routes } from "react-router-dom";
import { AdminRoute, AuthRoute, NonAuthRoute, PublicRoute } from "./middleware";

// pages layout
import PublicLayout from "./components/public/PublicLayout";
import AdminLayout from "./components/admin/AdminLayout";
import CustomerLayout from "./components/customer/CustomerLayout";
import ClassRecuring from "./pages/admin/classes/ClassRecuring";

function App() {
  const { checkingAuth, authMe } = useAuthStore();

  useEffect(() => {
    authMe();
  }, []);

  if (checkingAuth) return <Loading />;

  return (
    <>
      <Toaster />
      <ScrollToTop />
      <Routes>
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
        {/* Public */}
        <Route
          path="/"
          element={
            <PublicRoute>
              <PublicLayout />
            </PublicRoute>
          }
        >
          <Route index element={<Home />} />
          <Route path="about" element={<About />} />
          <Route path="classes" element={<Classes />} />
          <Route path="packages" element={<Packages />} />
          <Route path="schedules" element={<Schedules />} />
          <Route path="classes/:id" element={<ClassDetail />} />
          <Route path="packages/:id" element={<PackageDetail />} />
        </Route>

        {/* customer */}
        <Route
          path="/profile"
          element={
            <AuthRoute>
              <CustomerLayout />
            </AuthRoute>
          }
        >
          <Route index element={<Profile />} />
          <Route path="packages" element={<UserPackages />} />
          <Route path="bookings" element={<UserBookings />} />
          <Route path="transactions" element={<UserTransactions />} />
          <Route path="notifications" element={<UserNotifications />} />
        </Route>

        {/* admin */}
        <Route
          path="/admin"
          element={
            <AdminRoute>
              <AdminLayout />
            </AdminRoute>
          }
        >
          <Route path="users" element={<UsersList />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="classes" element={<ClassesList />} />
          <Route path="classes/add" element={<AddClass />} />
          <Route path="reviews" element={<ReviewsLists />} />
          <Route path="vouchers" element={<VouchersList />} />
          <Route path="bookings" element={<BookingsList />} />
          <Route path="packages" element={<PackagesList />} />

          <Route path="packages/add" element={<AddPackage />} />
          <Route path="instructors" element={<InstructorsList />} />
          <Route path="classes/options" element={<ClassOptions />} />
          <Route path="transactions" element={<TransactionsList />} />
          <Route path="instructors/add" element={<AddInstructors />} />
          <Route path="notifications" element={<NotificationsList />} />
          <Route path="classes/schedules" element={<ClassSchedules />} />
          <Route path="classes/recuring" element={<ClassRecuring />} />
          <Route index element={<Navigate to="dashboard" replace />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
