import React from "react";
import { Loading } from "@/components/ui/Loading";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { ProfileInfo } from "@/components/customer/profile/ProfileInfo";
import { UploadAvatar } from "@/components/customer/profile/UploadAvatar";
import { UpdateProfile } from "@/components/customer/profile/UpdateProfile";

const Profile = () => {
  const { data, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const profile = data || [];

  return (
    <section className="px-4 py-10 space-y-6">
      {/* Heading */}
      <div className="text-center space-y-1">
        <h2 className="text-2xl font-bold text-foreground">My Profile</h2>
        <p className="text-muted-foreground text-sm">
          Your personal account information
        </p>
      </div>

      {/* Card */}
      <div className="bg-card text-card-foreground border border-border rounded-xl shadow p-6 space-y-6">
        <div className="flex flex-col md:flex-row items-center gap-6">
          {/* Avatar */}
          <UploadAvatar profile={profile} />

          {/* Profile Info */}
          <ProfileInfo profile={profile} />
        </div>

        {/* Update Button */}
        <div className="flex justify-end">
          <UpdateProfile profile={profile} />
        </div>
      </div>
    </section>
  );
};

export default Profile;
