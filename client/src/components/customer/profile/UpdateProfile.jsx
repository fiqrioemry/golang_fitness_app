import React from "react";
import { Pencil } from "lucide-react";
import { profileSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormDialog } from "@/components/form/FormDialog";
import { SelectElement } from "@/components/input/SelectElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const UpdateProfile = ({ profile }) => {
  const { updateProfile } = useProfileMutation();
  return (
    <FormDialog
      state={profile}
      title="Edit Profil"
      schema={profileSchema}
      resourceId={profile.id}
      action={updateProfile.mutateAsync}
      buttonText={
        <Button size="sm">
          <Pencil className="w-4 h-4 mr-1" />
          Edit
        </Button>
      }
    >
      <InputTextElement name="fullname" label="Nama Lengkap" />
      <InputDateElement name="birthday" label="Tanggal Lahir" />
      <SelectElement
        name="gender"
        label="Jenis Kelamin"
        options={[
          { value: "male", label: "Laki-laki" },
          { value: "female", label: "Perempuan" },
        ]}
        placeholder="Pilih jenis kelamin"
        rules={{ required: false }}
      />
      <InputTextElement
        name="phone"
        label="Nomor Telepon"
        placeholder="08xxxx"
        isNumeric
        rules={{ required: false }}
      />
      <InputTextareaElement
        name="bio"
        label="Bio"
        placeholder="Ceritakan tentang dirimu..."
        rules={{ required: false }}
      />
    </FormDialog>
  );
};
