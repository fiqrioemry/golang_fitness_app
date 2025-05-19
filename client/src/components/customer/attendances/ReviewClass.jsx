import { Pencil } from "lucide-react";
import { Button } from "@/components/ui/button";
import { createReviewSchema } from "@/lib/schema";
import { useCreateReviewMutation } from "@/hooks/useReview";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputRatingElement } from "@/components/input/InputRatingElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const ReviewClass = ({ attendance }) => {
  const { mutate: createReview, isPending } = useCreateReviewMutation();
  const initialState = {
    classScheduleId: attendance.id,
    comment: "",
    rating: 0,
  };

  return (
    <FormAddDialog
      icon={true}
      state={initialState}
      schema={createReviewSchema}
      title="Create a comment"
      buttonText={
        <Button variant="outline">
          <Pencil />
          <span>Review Class</span>
        </Button>
      }
      loading={isPending}
      action={createReview}
    >
      <InputRatingElement name="rating" label="Your Rating" />
      <InputTextareaElement
        name="comment"
        label="Comment"
        placeholder="Write your comment here"
      />
    </FormAddDialog>
  );
};

export { ReviewClass };
