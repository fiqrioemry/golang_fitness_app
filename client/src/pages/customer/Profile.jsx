import React from "react";
import { Loading } from "@/components/ui/Loading";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { ProfileInfo } from "@/components/customer/profile/ProfileInfo";
import { UploadAvatar } from "@/components/customer/profile/UploadAvatar";
import { UpdateProfile } from "@/components/customer/profile/UpdateProfile";

const Profile = () => {
  const { data: profile, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-6xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Profile</h2>
        <p className="text-muted-foreground text-sm">
          Your personal account information
        </p>
      </div>

      <div className="bg-white shadow-sm border rounded-xl p-6">
        <div className="flex flex-col md:flex-row items-center gap-6">
          {/* avatar display */}
          <UploadAvatar profile={profile} />

          {/* profile display  */}
          <ProfileInfo profile={profile} />
        </div>

        {/* edit profile */}
        <div className="flex justify-end">
          <UpdateProfile profile={profile} />
        </div>
      </div>
    </section>
  );
};

export default Profile;
