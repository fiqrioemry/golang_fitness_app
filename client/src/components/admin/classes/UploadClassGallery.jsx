import React from "react";
import { uploadGallerySchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

const UploadClassGallery = ({ classes }) => {
  const { uploadGallery } = useClassMutation();
  const { isPending, mutateAsync } = uploadGallery;

  return (
    <FormAddDialog
      icon
      loading={isPending}
      schema={uploadGallerySchema}
      state={{ gallery: classes.galleries }}
      action={(data) => mutateAsync({ id: classes.id, gallery: data.gallery })}
      title="Upload Gallery"
    >
      <InputFileElement name="gallery" label="Gallery (Required)" />
    </FormAddDialog>
  );
};

export { UploadClassGallery };
