import {
  SidebarInset,
  SidebarTrigger,
  SidebarProvider,
} from "@/components/ui/sidebar";
import AdminSidebar from "./AdminSidebar";
import { Outlet } from "react-router-dom";
import { Separator } from "@/components/ui/separator";
import { MenuIcon } from "lucide-react";

const AdminLayout = () => {
  return (
    <SidebarProvider>
      <AdminSidebar />
      <SidebarInset>
        <header className="flex sticky top-0 bg-background h-16 shrink-0 items-center gap-2 border-b px-4 z-50">
          <SidebarTrigger className="-ml-1">
            <MenuIcon />
          </SidebarTrigger>
          <Separator orientation="vertical" className="mr-2 h-4" />
        </header>
        <div className="flex flex-1 flex-col bg-muted">
          <div className="py-3 md:py-6">
            <Outlet />
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
};

export default AdminLayout;
