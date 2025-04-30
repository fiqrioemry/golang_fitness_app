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

const CreateClass = () => {
  const { createClass } = useClassMutation();

  return (
    <section className="container mx-auto p-4">
      <FormInput
        className="w-full"
        state={createClassState}
        schema={createClassSchema}
        action={createClass.mutateAsync}
      >
        <InputTextElement
          name="title"
          label="Title kelas"
          placeholder="Masukkan nama kelas"
        />

        <div className="grid grid-cols-2 gap-4">
          <InputFileElement name="image" label="Thumbnail Kelas" isSingle />
          <div>
            <SelectElement
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
          <SelectElement
            name="categoryId"
            data="category"
            label="Kategori"
            placeholder="Pilih Kategori kelas"
          />
          <SelectElement
            name="levelId"
            data="level"
            label="Level"
            placeholder="Pilih Tingkat Kesulitan "
          />
          <SelectElement
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
      </FormInput>
    </section>
  );
};

export default CreateClass;
