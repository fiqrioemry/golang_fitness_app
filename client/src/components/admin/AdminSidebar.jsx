import {
  Plus,
  List,
  Bell,
  Clock,
  Users,
  Package,
  Dumbbell,
  UserPlus,
  BookUser,
  Calendar,
  BarChart2,
  LayoutList,
  ShoppingCart,
  TicketPercent,
  MessageSquare,
  ClipboardCheck,
  LogOut,
  Repeat,
} from "lucide-react";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "@/components/ui/accordion";

import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/sidebar";
import { cn } from "@/lib/utils";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";

const NavItem = ({ to, icon: Icon, title, active }) => (
  <Link
    to={to}
    className={cn(
      "flex items-center justify-between px-4 py-2 text-sm rounded-md transition",
      active ? "bg-blue-100 text-blue-600 font-semibold" : "hover:bg-gray-100"
    )}
  >
    <span className="flex items-center gap-2">
      <Icon className="w-4 h-4" />
      {title}
    </span>
  </Link>
);

const directMenus = [
  {
    to: "/admin/users",
    icon: Users,
    title: "Users",
  },
  {
    to: "/admin/bookings",
    icon: LayoutList,
    title: "Bookings",
  },
  {
    to: "/admin/transactions",
    icon: ShoppingCart,
    title: "Transactions",
  },
  {
    to: "/admin/vouchers",
    icon: TicketPercent,
    title: "Vouchers",
  },
  {
    to: "/admin/reviews",
    icon: MessageSquare,
    title: "Reviews",
  },
  {
    to: "/admin/notifications",
    icon: Bell,
    title: "Notifications",
  },
];

const accordionMenus = [
  {
    value: "instructor",
    icon: BookUser,
    title: "Instructors",
    children: [
      { to: "/admin/instructors", title: "Instructor List", icon: List },
      { to: "/admin/instructors/add", title: "Add Instructor", icon: UserPlus },
    ],
  },
  {
    value: "classes",
    icon: Dumbbell,
    title: "Classes",
    children: [
      { to: "/admin/classes", title: "Class List", icon: List },
      { to: "/admin/classes/add", title: "Add Class", icon: Plus },
      { to: "/admin/classes/options", title: "Class Options", icon: Calendar },
      { to: "/admin/classes/schedules", title: "Class Schedules", icon: Clock },
      {
        to: "/admin/classes/recuring",
        title: "Recuring Schedule",
        icon: Repeat,
      },
    ],
  },
  {
    value: "packages",
    icon: Package,
    title: "Packages",
    children: [
      { to: "/admin/packages", title: "Package List", icon: List },
      { to: "/admin/packages/add", title: "Add Package", icon: Plus },
    ],
  },
];

const AdminSidebar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const { user, logout } = useAuthStore();

  return (
    <Sidebar>
      <SidebarContent className="p-4 space-y-4 text-sm text-gray-700">
        <SidebarHeader className="mb-4">
          <h2>admin panel</h2>
        </SidebarHeader>

        <SidebarMenu className="space-y-1">
          <NavItem
            to="/admin"
            title="Dashboard"
            icon={BarChart2}
            active={currentPath === "/admin"}
          />

          {directMenus.map((item) => (
            <NavItem
              key={item.to}
              to={item.to}
              icon={item.icon}
              title={item.title}
              active={currentPath === item.to}
            />
          ))}

          <Accordion type="multiple" className="space-y-1">
            {accordionMenus.map((menu) => (
              <AccordionItem key={menu.value} value={menu.value}>
                <AccordionTrigger className="px-4 py-2 rounded-md hover:bg-gray-100 border-none">
                  <span className="flex items-center gap-2">
                    <menu.icon className="w-4 h-4" /> {menu.title}
                  </span>
                </AccordionTrigger>
                <AccordionContent className="pl-6 space-y-1 mt-1">
                  {menu.children.map((child) => (
                    <NavItem
                      key={child.to}
                      to={child.to}
                      title={child.title}
                      icon={child.icon}
                      active={currentPath === child.to}
                    />
                  ))}
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
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

export default AdminSidebar;
