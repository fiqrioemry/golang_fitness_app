import { LogIn } from "lucide-react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/store/useAuthStore";
import UserDropdown from "@/components/header/UserDropdown";
const Header = () => {
  const { user } = useAuthStore();
  const handleLoginClick = () => navigate("/signin");

  return (
    <div className="h-14 relative z-50">
      <header className="fixed w-full bg-white p-2 border-b shadow-sm">
        <div className="flex items-center justify-between container mx-auto gap-4">
          <Link to="/">
            <h2 className="text-xl font-bold text-primary">WELLNESS APP</h2>
          </Link>

          <div className="flex items-center gap-4">
            {/* User Dropdown Avatar & Login */}
            {user ? (
              <UserDropdown />
            ) : (
              <Button onClick={handleLoginClick}>
                <LogIn className="w-4 h-4" />
                Login
              </Button>
            )}
          </div>
        </div>
      </header>
    </div>
  );
};

export default Header;
