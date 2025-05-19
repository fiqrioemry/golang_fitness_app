// public pages
import Home from "./pages/Home";
import About from "./pages/About";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import Classes from "./pages/Classes";
import Packages from "./pages/Packages";
import NotFound from "./pages/NotFound";
import Schedules from "./pages/Schedules";
import ClassDetail from "./pages/ClassDetail";
import PackageDetail from "./pages/PackageDetail";
import ScheduleDetail from "./pages/ScheduleDetail";

// admin pages
import Dashboard from "./pages/admin/Dashboard";
import UsersList from "./pages/admin/UsersList";
import ClassAdd from "./pages/admin/classes/ClassAdd";
import Notifications from "./pages/admin/Notifications";
import PackageAdd from "./pages/admin/packages/PackageAdd";
import ClassesList from "./pages/admin/classes/ClassesList";
import InstructorsList from "./pages/admin/InstructorsList";
import VouchersAdd from "./pages/admin/vouchers/VouchersAdd";
import ClassOptions from "./pages/admin/classes/ClassOptions";
import TransactionsList from "./pages/admin/TransactionsList";
import VouchersList from "./pages/admin/vouchers/VouchersList";
import PackagesList from "./pages/admin/packages/PackagesList";
import ClassRecuring from "./pages/admin/classes/ClassRecuring";
import ClassSchedules from "./pages/admin/classes/ClassSchedules";

// customer pages
import Profile from "./pages/customer/Profile";
import UserBookings from "./pages/customer/UserBookings";
import UserSettings from "./pages/customer/UserSettings";
import UserPackages from "./pages/customer/UserPackages";
import UserAttendances from "./pages/customer/UserAttendances";
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
import AdminLayout from "./components/admin/AdminLayout";
import PublicLayout from "./components/public/PublicLayout";
import CustomerLayout from "./components/customer/CustomerLayout";

function App() {
  const { checkingAuth, authMe } = useAuthStore();

  useEffect(() => {
    authMe();
  }, []);

  if (checkingAuth) return <Loading />;

  return (
    <>
      <Toaster position="top-center" />
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
          <Route
            path="schedules/:id"
            element={
              <AuthRoute>
                <ScheduleDetail />
              </AuthRoute>
            }
          />
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
          <Route path="settings" element={<UserSettings />} />
          <Route path="attendances" element={<UserAttendances />} />
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
          <Route path="messages" element={<Notifications />} />
          <Route path="classes" element={<ClassesList />} />
          <Route path="classes/add" element={<ClassAdd />} />
          <Route path="classes/options" element={<ClassOptions />} />
          <Route path="classes/recuring" element={<ClassRecuring />} />
          <Route path="classes/schedules" element={<ClassSchedules />} />
          <Route path="vouchers" element={<VouchersList />} />
          <Route path="vouchers/add" element={<VouchersAdd />} />
          <Route path="packages" element={<PackagesList />} />
          <Route path="packages/add" element={<PackageAdd />} />
          <Route path="instructors" element={<InstructorsList />} />
          <Route path="transactions" element={<TransactionsList />} />
          <Route index element={<Navigate to="dashboard" replace />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
