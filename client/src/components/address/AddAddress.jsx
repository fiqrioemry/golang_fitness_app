// src/components/address/AddAddress.jsx
import React from "react";
import { PlusCircle } from "lucide-react";
import { addressSchema } from "@/lib/schema";
import { addressState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { useAddAddressMutation } from "@/hooks/useUserMutation";
import LocationSelection from "@/components/input/LocationSelection";

const AddAddress = () => {
  const { mutate: addAddress, isLoading } = useAddAddressMutation();

  return (
    <FormDialog
      loading={isLoading}
      action={addAddress}
      state={addressState}
      schema={addressSchema}
      title="Add New Address"
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>New Address</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Nama penerima"
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
        label="Alamat Penerima"
        isTextArea={true}
        placeholder="Masukkan Alamat Penerima"
      />

      <LocationSelection />

      <SwitchElement name="isMain" label="Atur sebagai alamat utama ?" />
    </FormDialog>
  );
};

export default AddAddress;
