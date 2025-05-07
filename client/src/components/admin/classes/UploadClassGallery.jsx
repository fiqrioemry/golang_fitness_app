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
      state={{ images: classes.galleries }}
      action={(data) => mutateAsync({ id: classes.id, images: data.images })}
      title="Upload Gallery"
    >
      <InputFileElement name="images" label="Gallery (Required)" />
    </FormAddDialog>
  );
};

export { UploadClassGallery };
