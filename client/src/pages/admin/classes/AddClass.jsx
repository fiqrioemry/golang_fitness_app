import React from "react";
import { createClassSchema } from "@/lib/schema";
import { createClassState } from "@/lib/constant";
import { useClassMutation } from "@/hooks/useClass";
import { FormInput } from "@/components/form/FormInput";
import { SelectElement } from "@/components/input/SelectElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

const AddClass = () => {
  const { createClass } = useClassMutation();

  return (
    <section className="max-w-5xl mx-auto px-6 py-6 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Tambah Kelas Baru</h2>
        <p className="text-muted-foreground text-sm">
          Isi form berikut untuk menambahkan kelas baru ke sistem.
        </p>
      </div>

      <div className="bg-white border shadow-sm rounded-xl p-6">
        <FormInput
          text="Add new class"
          className="space-y-6"
          state={createClassState}
          schema={createClassSchema}
          action={createClass.mutateAsync}
        >
          {/* Section 1: Judul & Thumbnail */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <InputTextElement
                name="title"
                label="Judul Kelas"
                placeholder="Masukkan nama kelas"
              />
              <InputTextareaElement
                name="description"
                label="Deskripsi"
                placeholder="Deskripsi minimal 20 karakter"
                maxLength={200}
              />
            </div>
            <InputFileElement
              name="image"
              label="Thumbnail Kelas"
              isSingle
              note="Rasio disarankan: 4:3"
            />
          </div>

          {/* Section 2: Lokasi, Durasi, Deskripsi */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <SelectElement
              label="Lokasi"
              data="location"
              name="locationId"
              placeholder="Pilih Lokasi kelas"
            />
            <InputNumberElement
              name="duration"
              label="Durasi (menit)"
              placeholder="Contoh: 60"
            />
          </div>

          {/* Section 3: Kategori, Level, Subkategori, Tipe */}
          <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
            <SelectElement
              name="categoryId"
              data="category"
              label="Kategori"
              placeholder="Pilih Kategori"
            />
            <SelectElement
              name="levelId"
              data="level"
              label="Level"
              placeholder="Tingkat Kesulitan"
            />
            <SelectElement
              name="subcategoryId"
              data="subcategory"
              label="Sub-Kategori"
              placeholder="Pilih Sub-Kategori"
            />
            <SelectOptionsElement
              name="typeId"
              data="type"
              label="Tipe Kelas"
              placeholder="Pilih Tipe"
            />
          </div>

          {/* Section 4: Informasi Tambahan & Gallery */}
          <div className="space-y-4">
            <InputTagsElement
              name="additional"
              label="Informasi Tambahan"
              placeholder="Tekan enter untuk tambah info"
            />
            <InputFileElement
              name="images"
              label="Galeri Tambahan (Opsional)"
              note="Bisa unggah lebih dari satu gambar"
            />
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default AddClass;
