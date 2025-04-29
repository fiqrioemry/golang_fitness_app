import { cn } from "@/lib/utils";
import { MapPin, Truck, UserRoundPen } from "lucide-react";
import { Link, Outlet, useLocation } from "react-router-dom";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";

const customerMenu = [
  {
    title: "profile",
    path: "/user/profile",
    icon: UserRoundPen,
  },
  {
    title: "address",
    path: "/user/address",
    icon: MapPin,
  },
  {
    title: "orders",
    path: "/user/orders",
    icon: Truck,
  },
];

const UserLayout = () => {
  const location = useLocation();

  const currentPath = location.pathname;

  return (
    <section className="section min-h-[62vh]">
      <Tabs defaultValue={currentPath}>
        <TabsList className="justify-between md:justify-start">
          {customerMenu.map((menu) => {
            const activePath = currentPath === menu.path;
            return (
              <TabsTrigger value={menu.path} key={menu.title} asChild>
                <Link
                  to={menu.path}
                  className={cn(
                    activePath ? "text-blue-700" : "text-foreground"
                  )}
                  key={menu.title}
                >
                  {menu.title}
                  <menu.icon className="ml-2 w-5 h-5 inline" />
                </Link>
              </TabsTrigger>
            );
          })}
        </TabsList>
      </Tabs>
      <div className="mt-4">{<Outlet />}</div>
    </section>
  );
};

export default UserLayout;
