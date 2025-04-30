import React from "react";
import { format } from "date-fns";
import UploadAvatar from "./UploadAvatar";
import UpdateProfile from "./UpdateProfile";
import { Loading } from "@/components/ui/Loading";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const Profile = () => {
  const { data: profile, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-4xl mx-auto px-6 py-10 space-y-8">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Profile</h2>
        <p className="text-sm text-muted-foreground">
          Informasi akun pribadi Anda
        </p>
      </div>

      <div className="bg-white shadow-sm border rounded-xl p-6">
        <div className="flex flex-col md:flex-row items-center gap-6">
          <div className="relative flex flex-col justify-center items-center space-y-4">
            <img
              src={profile.avatar}
              alt={profile.fullname}
              className="w-32 h-32 rounded-full object-cover border"
            />
            <UploadAvatar profile={profile} />
          </div>

          <div className="flex-1 space-y-3 w-full">
            <div className="flex justify-between items-start">
              <div>
                <h3 className="text-xl font-semibold">{profile.fullname}</h3>
                <p className="text-sm text-muted-foreground">{profile.email}</p>
              </div>
              <UpdateProfile profile={profile} />
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-muted-foreground mt-4">
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
        </div>
      </div>
    </section>
  );
};

export default Profile;
