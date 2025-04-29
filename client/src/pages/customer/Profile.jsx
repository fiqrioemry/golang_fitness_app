import React from "react";
import { useProfileQuery } from "@/hooks/useProfile";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";

const Profile = () => {
  const { data: profile, isError, refetch, isLoading } = useProfileQuery();
  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return <div></div>;
};

export default Profile;
