// src/components/address/UpdateClass.jsx
import React from "react";
import { createClassSchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { FormDialog } from "@/components/form/FormDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { Pencil } from "lucide-react";

const UpdateClass = ({ classes }) => {
  const { updateClass, isLoading } = useClassMutation();

  return (
    <FormDialog
      loading={isLoading}
      resourceId={classes.id}
      title="Perbaharui Kelas"
      state={classes}
      schema={createClassSchema}
      action={updateClass.mutateAsync}
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
        name="title"
        label="Title kelas"
        placeholder="Masukkan nama kelas"
      />
      <div className="grid grid-cols-2 gap-4">
        <InputFileElement name="image" label="Thumbnail Kelas" isSingle />
        <div>
          <SelectOptionsElement
            label="Lokasi"
            data="location"
            name="locationId"
            placeholder="Pilih Lokasi kelas "
          />
          <InputNumberElement
            name="duration"
            label="Durasi"
            placeholder="Durasi waktu kelas dalam menit"
          />

          <InputTextareaElement
            maxLength={200}
            name="description"
            label="Deskripsi kelas"
            placeholder="Masukkan deskripsi kelas min. 20 karakter"
          />
        </div>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <SelectOptionsElement
          name="categoryId"
          data="category"
          label="Kategori"
          placeholder="Pilih Kategori kelas"
        />
        <SelectOptionsElement
          name="levelId"
          data="level"
          label="Level"
          placeholder="Pilih Tingkat Kesulitan "
        />
        <SelectOptionsElement
          name="subcategoryId"
          data="subcategory"
          label="Sub-Kategori"
          placeholder="Pilih Sub-Kategori kelas "
        />
        <SelectOptionsElement
          data="type"
          name="typeId"
          label="Tipe kelas"
          placeholder="Pilih Tipe Kelas "
        />
      </div>
      <InputTagsElement
        name="additional"
        label="Informasi tambahan"
        placeholder="Masukkan informasi, tekan enter untuk tambah"
      />
      <InputFileElement name="images" label="Gallery (Optional)" />
      <SwitchElement name="isActive" label="Atur sebagai kelas aktif" />
    </FormDialog>
  );
};

export default UpdateClass;
