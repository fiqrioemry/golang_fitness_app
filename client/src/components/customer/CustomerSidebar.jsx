import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/sidebar";
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
  Settings2,
} from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";

const NavItem = ({ to, icon: Icon, title, active }) => (
  <Link
    to={to}
    className={cn(
      "flex items-center gap-3 px-4 py-2 text-sm rounded-md transition",
      active
        ? "bg-primary text-primary-foreground font-semibold"
        : "text-muted-foreground hover:bg-muted"
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
      <SidebarContent className="p-4 space-y-4 text-sm">
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
            icon={Calendar}
            title="My Attendance"
            to="/profile/attendances"
            active={currentPath === "/profile/attendances"}
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
          <NavItem
            to="/profile/settings"
            title="Settings"
            icon={Settings2}
            active={currentPath === "/profile/settings"}
          />
        </SidebarMenu>
      </SidebarContent>

      <SidebarFooter className="p-4 text-xs text-muted-foreground">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3 cursor-pointer hover:bg-muted px-3 py-2 rounded-md transition">
              <Avatar className="w-9 h-9">
                <AvatarImage src={user?.avatar} alt={user?.fullname} />
                <AvatarFallback>{user?.fullname?.[0] || "A"}</AvatarFallback>
              </Avatar>
              <div className="flex flex-col text-left overflow-hidden">
                <span className="text-sm font-medium text-foreground truncate">
                  {user?.fullname || "Admin"}
                </span>
                <span className="text-xs text-muted-foreground truncate">
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
