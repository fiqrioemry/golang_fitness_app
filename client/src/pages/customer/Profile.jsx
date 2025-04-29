import React from "react";
import { Loading } from "@/components/ui/Loading";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const Profile = () => {
  const { data: profile, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return <div></div>;
};

export default Profile;
