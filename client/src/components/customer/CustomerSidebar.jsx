import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/sidebar";
import { cn } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import {
  Calendar,
  Package,
  ShoppingCart,
  Bell,
  User,
  LogOut,
} from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

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
  const location = useLocation();
  const currentPath = location.pathname;
  const { user, logout } = useAuthStore();

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
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3 cursor-pointer hover:bg-gray-100 p-2 rounded-md transition">
              <Avatar className="w-9 h-9">
                <AvatarImage src={user?.avatar} alt={user?.fullname} />
                <AvatarFallback>{user?.fullname?.[0] || "A"}</AvatarFallback>
              </Avatar>
              <div className="flex flex-col text-left">
                <span className="text-sm font-medium text-gray-700">
                  {user?.fullname || "Admin"}
                </span>
                <span className="text-xs text-muted-foreground truncate max-w-[140px]">
                  {user?.email || "admin@gmail.com"}
                </span>
              </div>
            </div>
          </DropdownMenuTrigger>

          <DropdownMenuContent side="top" align="start" className="w-60">
            <DropdownMenuItem onClick={logout}>
              <LogOut className="w-4 h-4 mr-2" />
              Logout
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarFooter>
    </Sidebar>
  );
};

export default CustomerSidebar;
