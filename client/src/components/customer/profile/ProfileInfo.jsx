import React from "react";
import { format } from "date-fns";

export const ProfileInfo = ({ profile }) => {
  return (
    <div className="flex-1 space-y-3 w-full">
      <div className="flex justify-between items-start">
        <div>
          <h3 className="text-xl font-semibold">{profile.fullname}</h3>
          <p className="text-sm text-muted-foreground">{profile.email}</p>
        </div>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-muted-foreground">
        <p>
          <span className="font-medium text-gray-800">Phone:</span>{" "}
          {profile.phone || "-"}
        </p>
        <p>
          <span className="font-medium text-gray-800">Gender:</span>{" "}
          {profile.gender || "-"}
        </p>
        <p>
          <span className="font-medium text-gray-800">Birthday:</span>{" "}
          {profile.birthday
            ? format(new Date(profile.birthday), "dd MMMM yyyy")
            : "-"}
        </p>
        <p>
          <span className="font-medium text-gray-800">Joined:</span>{" "}
          {format(new Date(profile.joinedAt), "dd MMMM yyyy")}
        </p>
      </div>

      <div className="mt-4 text-sm text-muted-foreground">
        <p className="font-medium text-gray-800">Bio:</p>
        <p>{profile.bio || "-"}</p>
      </div>
    </div>
  );
};
