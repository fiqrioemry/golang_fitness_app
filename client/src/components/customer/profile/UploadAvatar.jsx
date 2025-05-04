import React from "react";
import { avatarSchema } from "@/lib/schema";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

export const UploadAvatar = ({ profile }) => {
  const { updateAvatar } = useProfileMutation();
  const { mutateAsync, isPending } = updateAvatar;

  return (
    <div className="relative flex flex-col justify-center items-center space-y-4">
      <img
        src={profile.avatar}
        alt={profile.fullname}
        className="w-32 h-32 rounded-full object-cover border"
      />
      <FormUpdateDialog
        state={profile}
        title="Edit Avatar"
        schema={avatarSchema}
        loading={isPending}
        action={mutateAsync}
      >
        <InputFileElement
          isSingle
          name="avatar"
          label="Upload Avatar"
          note="Rekomendasi: rasio 1:1 (400x400px)"
        />
      </FormUpdateDialog>
    </div>
  );
};
