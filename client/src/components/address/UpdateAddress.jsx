// src/components/address/UpdateAddress.jsx
import React from "react";
import { Edit2 } from "lucide-react";
import { addressSchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { useUpdateAddressMutation } from "@/hooks/useUserMutation";
import LocationSelection from "@/components/input/LocationSelection";

const UpdateAddress = ({ address }) => {
  const { mutate: updateAddress, isLoading } = useUpdateAddressMutation();

  return (
    <FormDialog
      state={address}
      loading={isLoading}
      action={updateAddress}
      resourceId={address.id}
      schema={addressSchema}
      title="Update Address"
      buttonText={
        <button className="btn btn-secondary">
          <Edit2 className="text-white-500" size={18} />
        </button>
      }
    >
      <InputElement
        name="name"
        label="Nama Penerima"
        placeholder="Masukkan nama penerima"
      />
      <InputElement
        name="phone"
        maxLength={12}
        isNumeric={true}
        label="Nomor Telepon"
        placeholder="Masukkan nomor telepon"
      />
      <InputElement
        name="address"
        isTextArea
        label="Alamat Lengkap"
        placeholder="Masukkan alamat lengkap"
      />

      <LocationSelection />
      <SwitchElement name="isMain" label="Atur sebagai alamat utama?" />
    </FormDialog>
  );
};

export default UpdateAddress;
