import React from "react";
import { Camera } from "lucide-react";
import { avatarSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormDialog } from "@/components/form/FormDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

const UploadAvatar = ({ profile }) => {
  const { updateAvatar } = useProfileMutation();

  return (
    <FormDialog
      state={profile}
      title="Edit Avatar"
      schema={avatarSchema}
      action={updateAvatar.mutateAsync}
      buttonText={
        <Button size="sm">
          <Camera />
          Upload Avatar
        </Button>
      }
    >
      <InputFileElement
        isSingle
        name="avatar"
        label="Upload Avatar"
        note="Rekomendasi: rasio 1:1 (400x400px)"
      />
    </FormDialog>
  );
};

export default UploadAvatar;
