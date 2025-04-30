import { NavLink } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/store/useAuthStore";
import UserDropdown from "@/components/header/UserDropdown";

const Header = () => {
  const { user } = useAuthStore();
  const linkClass =
    "hover:text-blue-600 font-medium transition-colors duration-200";
  const activeClass =
    "text-blue-600 font-semibold underline underline-offset-4";

  return (
    <header className="fixed top-0 w-full bg-white/70 backdrop-blur-md shadow z-50">
      <div className="max-w-7xl mx-auto px-4 py-3 flex justify-between items-center">
        <NavLink to="/" className="text-2xl font-bold text-gray-800">
          FitBook
        </NavLink>
        <div className="hidden md:flex gap-6 items-center">
          <NavLink
            to="/"
            end
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Home
          </NavLink>
          <NavLink
            to="/classes"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Class
          </NavLink>
          <NavLink
            to="/packages"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Package
          </NavLink>
          <NavLink
            to="/schedule"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Schedule
          </NavLink>
          <NavLink
            to="/location"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Locate Us
          </NavLink>
          {user ? (
            <UserDropdown />
          ) : (
            <NavLink to="/signin">
              <Button className="px-5 py-2">Get Started</Button>
            </NavLink>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
