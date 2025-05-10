import React from "react";
import { PencilIcon } from "lucide-react";
import { profileSchema } from "@/lib/schema";
import { genderOptions } from "@/lib/constant";
import { Button } from "@/components/ui/button";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { SelectElement } from "@/components/input/SelectElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const UpdateProfile = ({ profile }) => {
  const { updateProfile } = useProfileMutation();

  return (
    <FormAddDialog
      icon={true}
      state={profile}
      title="Edit Profile"
      loading={updateProfile.isPending}
      action={updateProfile.mutateAsync}
      schema={profileSchema}
      buttonText={
        <Button variant="secondary" type="button">
          <PencilIcon className="w-4 h-4" />
          <span>Edit Profile</span>
        </Button>
      }
    >
      <InputTextElement name="fullname" label="Full Name" />
      <InputDateElement name="birthday" label="Date of Birth" />
      <SelectElement
        name="gender"
        label="Gender"
        options={genderOptions}
        placeholder="Select gender"
        rules={{ required: false }}
      />
      <InputTextElement
        name="phone"
        label="Phone Number"
        placeholder="08xxxx"
        isNumeric
        rules={{ required: false }}
      />
      <InputTextareaElement
        name="bio"
        label="Bio"
        placeholder="Tell us about yourself..."
        rules={{ required: false }}
      />
    </FormAddDialog>
  );
};
