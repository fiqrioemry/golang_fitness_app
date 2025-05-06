import React from "react";
import { format } from "date-fns";

export const ProfileInfo = ({ profile }) => {
  return (
    <div className="flex-1 w-full space-y-4 text-sm text-muted-foreground">
      {/* Header */}
      <div className="flex justify-between items-start">
        <div>
          <h2 className="text-xl font-semibold text-foreground">
            {profile.fullname}
          </h2>
          <p className="text-muted-foreground">{profile.email}</p>
        </div>
      </div>

      {/* Detail Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <p>
          <span className="font-medium text-foreground">Phone:</span>{" "}
          {profile.phone || "-"}
        </p>
        <p>
          <span className="font-medium text-foreground">Gender:</span>{" "}
          {profile.gender || "-"}
        </p>
        <p>
          <span className="font-medium text-foreground">Birthday:</span>{" "}
          {profile.birthday
            ? format(new Date(profile.birthday), "dd MMMM yyyy")
            : "-"}
        </p>
        <p>
          <span className="font-medium text-foreground">Joined:</span>{" "}
          {format(new Date(profile.joinedAt), "dd MMMM yyyy")}
        </p>
      </div>

      {/* Bio */}
      <div className="pt-2">
        <p className="font-medium text-foreground">Bio:</p>
        <p>{profile.bio || "-"}</p>
      </div>
    </div>
  );
};
