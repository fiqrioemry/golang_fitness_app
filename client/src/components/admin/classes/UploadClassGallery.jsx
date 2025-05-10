import { uploadGallerySchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputFileElement } from "@/components/input/InputFileElement";
import { CirclePlus } from "lucide-react";
import { Button } from "@/components//ui/button";

const UploadClassGallery = ({ classes }) => {
  const { uploadGallery } = useClassMutation();

  return (
    <FormAddDialog
      icon={true}
      title="Upload Gallery"
      schema={uploadGallerySchema}
      loading={uploadGallery.isPending}
      state={{ images: classes.galleries }}
      buttonText={
        <Button size="icon" variant="secondary">
          <CirclePlus />
        </Button>
      }
      action={(data) =>
        uploadGallery.mutateAsync({ id: classes.id, images: data.images })
      }
    >
      <InputFileElement name="images" label="Gallery (Required)" />
    </FormAddDialog>
  );
};

export { UploadClassGallery };
