// src/components/header/UserDropdown.jsx

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/useAuthStore";
import { LogOut, User, MapPin } from "lucide-react";

const UserDropdown = () => {
  const navigate = useNavigate();
  const { user, logout } = useAuthStore();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <img
          src={user.profile?.avatar}
          alt={user.profile?.fullname}
          className="w-8 h-8 rounded-full object-cover cursor-pointer"
        />
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="w-48 shadow-lg rounded-xl">
        <DropdownMenuItem
          onClick={() => navigate("/user/profile")}
          className="cursor-pointer"
        >
          <User className="w-4 h-4 mr-2" />
          Profile
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() => navigate("/user/address")}
          className="cursor-pointer"
        >
          <MapPin className="w-4 h-4 mr-2" />
          Address
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() => logout()}
          className="cursor-pointer text-red-500"
        >
          <LogOut className="w-4 h-4 mr-2" />
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UserDropdown;
