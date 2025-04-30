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
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "@/components/ui/accordion";
import {
  BarChart2,
  Users,
  Dumbbell,
  Clock,
  Package,
  ShoppingCart,
  UserPlus,
  CalendarClock,
  Plus,
  List,
  Bell,
  LayoutList,
  TicketPercent,
  BookUser,
  MessageSquare,
  ClipboardCheck,
  Calendar,
} from "lucide-react";

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

const accordionMenus = [
  {
    value: "user",
    icon: Users,
    title: "Users",
    children: [{ to: "/admin/users", title: "User List", icon: List }],
  },
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
  {
    value: "booking",
    icon: LayoutList,
    title: "Booking",
    children: [
      { to: "/admin/bookings", title: "All Bookings", icon: ClipboardCheck },
    ],
  },
  {
    value: "transaction",
    icon: ShoppingCart,
    title: "Transactions",
    children: [
      { to: "/admin/transactions", title: "Payments", icon: ShoppingCart },
    ],
  },
  {
    value: "voucher",
    icon: TicketPercent,
    title: "Vouchers",
    children: [{ to: "/admin/vouchers", title: "Voucher List", icon: List }],
  },
  {
    value: "review",
    icon: MessageSquare,
    title: "Reviews",
    children: [{ to: "/admin/reviews", title: "Review List", icon: List }],
  },
  {
    value: "notification",
    icon: Bell,
    title: "Notifications",
    children: [
      { to: "/admin/notifications", title: "Notification List", icon: List },
    ],
  },
];

const AdminSidebar = () => {
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
            to="/admin"
            title="Dashboard"
            icon={BarChart2}
            active={currentPath === "/admin"}
          />

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
        <div>
          Logged in as <strong>{user?.fullname || "Admin"}</strong>
        </div>
        <div className="truncate">{user?.email || "admin@gmail.com"}</div>
      </SidebarFooter>
    </Sidebar>
  );
};

export default AdminSidebar;
