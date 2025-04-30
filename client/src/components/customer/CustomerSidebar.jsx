import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/sidebar";
import { cn } from "@/lib/utils";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import { Calendar, Package, ShoppingCart, Bell, User } from "lucide-react";

const NavItem = ({ to, icon: Icon, title, active }) => (
  <Link
    to={to}
    className={cn(
      "flex items-center gap-2 px-4 py-2 text-sm rounded-md transition",
      active ? "bg-blue-100 text-blue-600 font-semibold" : "hover:bg-gray-100"
    )}
  >
    <Icon className="w-4 h-4" />
    {title}
  </Link>
);

const CustomerSidebar = () => {
  const { user } = useAuthStore();
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <Sidebar>
      <SidebarContent className="p-4 space-y-4 text-sm text-gray-700">
        <SidebarHeader className="mb-4">
          <WebLogo />
        </SidebarHeader>

        <SidebarMenu className="space-y-1">
          <NavItem
            to="/profile"
            title="My Profile"
            icon={User}
            active={currentPath === "/profile"}
          />
          <NavItem
            to="/profile/packages"
            title="My Packages"
            icon={Package}
            active={currentPath === "/profile/packages"}
          />
          <NavItem
            to="/profile/bookings"
            title="My Bookings"
            icon={Calendar}
            active={currentPath === "/profile/bookings"}
          />
          <NavItem
            to="/profile/transactions"
            title="Transaction History"
            icon={ShoppingCart}
            active={currentPath === "/profile/transactions"}
          />
          <NavItem
            to="/profile/notifications"
            title="Notifications"
            icon={Bell}
            active={currentPath === "/profile/notifications"}
          />
        </SidebarMenu>
      </SidebarContent>

      <SidebarFooter className="p-4 text-xs text-gray-500">
        <div>
          Logged in as <strong>{user?.fullname || "Customer"}</strong>
        </div>
        <div className="truncate">{user?.email || "customer@example.com"}</div>
      </SidebarFooter>
    </Sidebar>
  );
};

export default CustomerSidebar;
