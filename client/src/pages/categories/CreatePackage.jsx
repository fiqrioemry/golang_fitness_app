import React from "react";
import { useNavigate } from "react-router-dom";
import { packageState } from "@/lib/constant";
import { usePackageMutation } from "@/hooks/usePackage";
import { FormInput } from "@/components/form/FormInput";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextAreElement";
const CreatePackage = () => {
  const navigate = useNavigate();
  const { createPackage } = usePackageMutation();

  return (
    <section className="container mx-auto p-6 space-y-6">
      <div className="flex items-center justify-between">
        <h3 className="text-2xl font-bold">Create New Package</h3>
        <button
          onClick={() => navigate("/admin/packages")}
          className="text-primary hover:underline text-sm"
        >
          ← Kembali ke daftar paket
        </button>
      </div>

      <div className="bg-white rounded-md shadow-sm">
        <FormInput
          className="w-full space-y-6"
          state={packageState}
          schema={packageState}
          action={createPackage.mutateAsync}
        >
          {/* Baris Nama & Deskripsi */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
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
              {/* Baris Harga, Credit, Expired */}
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <InputNumberElement
                  name="price"
                  label="Harga Paket (Rp)"
                  placeholder="Masukkan harga dalam satuan Rp"
                />
                <InputNumberElement
                  name="credit"
                  label="Jumlah Credit"
                  placeholder="Jumlah credit dalam unit sesi"
                />
                <InputNumberElement
                  name="expired"
                  label="Expired (Hari)"
                  placeholder="Durasi masa aktif paket"
                />
              </div>
            </div>

            {/* Upload Thumbnail */}
            <InputFileElement name="image" label="Thumbnail Gambar" isSingle />
          </div>

          {/* Informasi Tambahan */}
          <InputTagsElement
            name="information"
            label="Informasi Tambahan Paket"
            placeholder="Masukkan informasi, tekan Enter untuk tambah"
          />

          {/* Status Aktif */}
          <SwitchElement name="isActive" label="Aktifkan Paket Ini" />
        </FormInput>
      </div>
    </section>
  );
};

export default CreatePackage;
