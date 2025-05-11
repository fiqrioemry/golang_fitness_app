import { usePackageMutation } from "@/hooks/usePackage";
import { FormInput } from "@/components/form/FormInput";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectElement } from "@/components/input/SelectElement";
import { notificationState, typeCode } from "@/lib/constant";
import { notificationSchema } from "@/lib/schema";
import { useSendPromoNotification } from "@/hooks/useNotification";

const Notifications = () => {
  const { mutate: sendPromo, isPending } = useSendPromoNotification();

  return (
    <section className="section">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Send New Message</h2>
        <p className="text-muted-foreground text-sm">
          Send a notification through all your users with just on click
        </p>
      </div>

      <div className="bg-background  rounded-xl shadow-sm border p-6">
        <FormInput
          className="w-full md:w-72"
          state={notificationState}
          schema={notificationSchema}
          text={"Send Promo Message"}
          isLoading={isPending}
          action={sendPromo}
        >
          <div className="space-y-4">
            <InputTextElement
              name="title"
              label="Notification Title"
              placeholder="Enter the title for notifications"
            />
            <InputTextareaElement
              name="message"
              label="Notification Message"
              placeholder="Write a message for notifications"
              maxLength={200}
            />
            <SelectElement
              name="typeCode"
              options={typeCode}
              label="Notification Type"
              placeholder="Select Notification Type"
            />
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default Notifications;
