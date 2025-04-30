// src/components/address/UpdateClass.jsx
import React from "react";
import { Pencil } from "lucide-react";
import { packageSchema } from "@/lib/schema";
import { usePackageMutation } from "@/hooks/usePackage";
import { FormDialog } from "@/components/form/FormDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextAreElement";

const UpdatePackage = ({ pkg }) => {
  const { updatePackage, isLoading } = usePackageMutation();

  return (
    <FormDialog
      state={pkg}
      loading={isLoading}
      resourceId={pkg.id}
      title="Perbaharui Package"
      schema={packageSchema}
      action={updatePackage.mutateAsync}
      buttonText={
        <button
          type="button"
          className="text-primary hover:text-blue-600 transition"
        >
          <Pencil className="w-4 h-4" />
        </button>
      }
    >
      <InputTextElement
        name="name"
        label="Nama Paket"
        placeholder="Masukkan nama paket"
      />
      <InputTextareaElement
        maxLength={200}
        name="description"
        label="Deskripsi Paket"
        placeholder="Masukkan deskripsi paket min. 20 karakter"
      />

      <InputNumberElement
        name="price"
        label="Harga paket"
        placeholder="Harga paket dalam satuan Rp"
      />
      <InputNumberElement
        name="credit"
        label="Credit"
        placeholder="Jumlah credit dalam satuan Unit"
      />
      <InputNumberElement
        name="expired"
        label="Waktu Expired"
        placeholder="Durasi waktu expired paket dalam hari"
      />
      <InputTagsElement
        name="information"
        label="Informasi paket"
        placeholder="Masukkan informasi, tekan enter untuk tambah"
      />
      <InputFileElement name="image" label="Gambar Thumbnail" isSingle />
      <SwitchElement name="isActive" label="Atur sebagai Paket aktif" />
    </FormDialog>
  );
};

export default UpdatePackage;
