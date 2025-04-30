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
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const AddPackage = () => {
  const { createPackage } = usePackageMutation();

  return (
    <section className="max-w-5xl mx-auto px-6 py-10 space-y-8">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Tambah Paket Baru</h2>
        <p className="text-muted-foreground text-sm">
          Lengkapi informasi paket yang akan ditawarkan kepada pelanggan.
        </p>
      </div>

      <div className="bg-white rounded-xl shadow-sm border p-6">
        <FormInput
          className="space-y-6"
          state={packageState}
          schema={packageState}
          action={createPackage.mutateAsync}
        >
          {/* Informasi Utama */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="space-y-4">
              <InputTextElement
                name="name"
                label="Nama Paket"
                placeholder="Masukkan nama paket"
              />
              <InputTextareaElement
                name="description"
                label="Deskripsi Paket"
                placeholder="Deskripsi minimal 20 karakter"
                maxLength={200}
              />

              <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
                <InputNumberElement
                  name="price"
                  label="Harga (Rp)"
                  placeholder="Contoh: 500000"
                />
                <InputNumberElement
                  name="credit"
                  label="Jumlah Credit"
                  placeholder="Contoh: 5"
                />
                <InputNumberElement
                  name="expired"
                  label="Expired (hari)"
                  placeholder="Contoh: 60"
                />
              </div>

              <InputTagsElement
                name="information"
                label="Informasi Tambahan"
                placeholder="Contoh: Tidak dapat direfund, tekan Enter"
              />

              <SwitchElement
                name="isActive"
                label="Status Aktif"
                description="Aktifkan agar paket tampil di halaman pembelian user."
              />
            </div>

            <div>
              <InputFileElement
                isSingle
                name="image"
                label="Thumbnail Paket"
                note="Rekomendasi: rasio 1:1 (400x400px)"
              />
            </div>
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default AddPackage;
