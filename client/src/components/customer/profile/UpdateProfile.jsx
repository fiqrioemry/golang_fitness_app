import React from "react";
import { profileSchema } from "@/lib/schema";
import { useProfileMutation } from "@/hooks/useProfile";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const UpdateProfile = ({ profile }) => {
  const { updateProfile } = useProfileMutation();
  const { mutateAsync, isPending } = updateProfile;
  return (
    <FormUpdateDialog
      icon={false}
      state={profile}
      title="Edit Profil"
      loading={isPending}
      action={mutateAsync}
      schema={profileSchema}
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
    </FormUpdateDialog>
  );
};
