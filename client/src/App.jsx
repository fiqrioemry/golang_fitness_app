// src/App.jsx
// public pages
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import Classes from "./pages/Classes";
import NotFound from "./pages/NotFound";
import ClassDetail from "./pages/ClassDetail";
import Profile from "./pages/customer/Profile";

// admin pages
import Dashboard from "./pages/admin/Dashboard";
import UsersList from "./pages/admin/UsersList";
import AddClass from "./pages/admin/classes/AddClass";
import BookingsList from "./pages/admin/BookingsList";
import VouchersList from "./pages/admin/VouchersList";
import ReviewsLists from "./pages/admin/ReviewsLists";
import Notifications from "./pages/admin/Notifications";
import AddPackage from "./pages/admin/packages/AddPackage";
import ClassesList from "./pages/admin/classes/ClassesList";
import ClassOptions from "./pages/admin/classes/ClassOptions";
import TransactionsList from "./pages/admin/TransactionsList";
import PackagesList from "./pages/admin/packages/PackagesList";
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
import Layout from "./components/layout/Layout";
import ScrollToTop from "./hooks/useScrollToTop";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "./store/useAuthStore";
import { AuthRoute, NonAuthRoute } from "./middleware";
import { Navigate, Route, Routes } from "react-router-dom";
import AdminLayout from "./components/dashboard/AdminLayout";
import CustomerLayout from "./components/customer/CustomerLayout";

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
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="classes" element={<Classes />} />
          <Route path="packages" element={<Packages />} />
          <Route path="packages/:id" element={<PackageDetail />} />
          <Route path="classes/:id" element={<ClassDetail />} />
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
        <Route path="/admin" element={<AdminLayout />}>
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="classes" element={<ClassesList />} />
          <Route path="classes/add" element={<AddClass />} />
          <Route path="classes/options" element={<ClassOptions />} />
          <Route path="classes/schedules" element={<ClassSchedules />} />
          <Route path="instructors" element={<InstructorsList />} />
          <Route path="instructors/add" element={<AddInstructors />} />
          <Route path="users" element={<UsersList />} />
          <Route path="reviews" element={<ReviewsLists />} />
          <Route path="vouchers" element={<VouchersList />} />
          <Route path="bookings" element={<BookingsList />} />
          <Route path="packages" element={<PackagesList />} />
          <Route path="packages/add" element={<AddPackage />} />
          <Route path="notifications" element={<Notifications />} />
          <Route path="transactions" element={<TransactionsList />} />
          <Route index element={<Navigate to="dashboard" replace />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
