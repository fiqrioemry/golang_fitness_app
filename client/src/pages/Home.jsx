import React from "react";
import { createClassSchema } from "@/lib/schema";
import { createClassState } from "@/lib/constant";
import { useLevelsQuery } from "@/hooks/useLevel";
import { Loading } from "@/components/ui/Loading";
import { useClassMutation } from "@/hooks/useClass";
import ErrorDialog from "@/components/ui/ErrorDialog";
import { FormInput } from "@/components/form/FormInput";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { SelectElement } from "@/components/input/SelectElement";
import { useSubcategoriesQuery } from "@/hooks/useSubcategories";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextAreElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { useLocationsQuery } from "@/hooks/useLocation";
import { InputImageElement } from "../components/input/InputImageElement";

const Home = () => {
  const { data: locations } = useLocationsQuery();
  const { data: categories } = useCategoriesQuery();
  const { mutate: createClass } = useClassMutation();
  const { data: subcategories } = useSubcategoriesQuery();
  const { data: levels, isLoading, isError } = useLevelsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="container mx-auto p-4">
      <FormInput
        className="w-full"
        action={createClass}
        state={createClassState}
        schema={createClassSchema}
      >
        <InputTextElement
          name="title"
          label="Title kelas"
          placeholder="Masukkan nama kelas"
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
        <InputTagsElement
          name="additional"
          label="Informasi tambahan"
          placeholder="Masukkan informasi, tekan enter untuk tambah"
        />

        <div className="grid grid-cols-2 gap-4">
          <InputImageElement name="image" label="Thumbnail Kelas" isSingle />
          <div className="space-y-3">
            <SelectElement
              name="categoryId"
              options={categories}
              label="Kategori"
              placeholder="Pilih Kategori kelas"
            />
            <SelectElement
              name="levelId"
              options={levels}
              label="Level"
              placeholder="Pilih Tingkat Kesulitan "
            />
            <SelectElement
              name="subcategoryId"
              options={subcategories}
              label="Sub-Kategori"
              placeholder="Pilih Sub-Kategori kelas "
            />
            <SelectElement
              name="locationId"
              options={locations}
              label="Lokasi"
              placeholder="Pilih Lokasi kelas "
            />
          </div>
        </div>
        <InputFileElement name="images" label="Gallery Kelas (Optional)" />
      </FormInput>
    </section>
  );
};

export default Home;
