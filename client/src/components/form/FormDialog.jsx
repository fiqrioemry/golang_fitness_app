import {
  Dialog,
  DialogTitle,
  DialogTrigger,
  DialogContent,
  DialogDescription,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { ScrollArea } from "@/components/ui/scroll-area";
import { SubmitLoading } from "@/components/ui/SubmitLoading";
import { SubmitButton } from "@/components/form/SubmitButton";
import { useState, useCallback, useMemo, useEffect } from "react";

export function FormDialog({
  title,
  state,
  schema,
  action,
  children,
  resourceId = null,
  loading = false,
  shouldReset = true,
  buttonText = "Update",
}) {
  const [isOpen, setIsOpen] = useState(false);
  const [showConfirmation, setShowConfirmation] = useState(false);

  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
    mode: "onChange",
  });

  const { formState, reset, handleSubmit } = methods;

  const isFormDirty = useMemo(() => formState.isDirty, [formState.isDirty]);

  const resetAndCloseDialog = useCallback(() => {
    reset();
    setIsOpen(false);
  }, [reset]);

  const handleCancel = useCallback(() => {
    if (isFormDirty) setShowConfirmation(true);
    else resetAndCloseDialog();
  }, [isFormDirty, resetAndCloseDialog]);

  const handleConfirmation = useCallback(
    (confirmed) => {
      if (confirmed) resetAndCloseDialog();
      setShowConfirmation(false);
    },
    [resetAndCloseDialog]
  );

  useEffect(() => {
    if (state) reset(state);
  }, [state, reset]);

  const handleSave = useCallback(
    async (data) => {
      if (resourceId !== null && resourceId !== undefined) {
        await action({ id: resourceId, data });
      } else {
        await action(data);
      }
      if (formState.isValid && shouldReset) reset();
      setIsOpen(false);
    },
    [action, formState.isValid, reset, shouldReset, resourceId]
  );

  return (
    <>
      {/* Main Dialog */}
      <Dialog
        open={isOpen}
        onOpenChange={(open) => (!open ? handleCancel() : setIsOpen(open))}
      >
        <DialogTrigger asChild>{buttonText}</DialogTrigger>

        <DialogContent className="sm:max-w-lg overflow-hidden rounded-xl p-0">
          {loading ? (
            <SubmitLoading />
          ) : (
            <FormProvider {...methods}>
              <form
                onSubmit={handleSubmit(handleSave)}
                className="flex flex-col h-[65vh]"
              >
                {/* Header */}
                <div className="border-b px-6 py-4">
                  <DialogTitle className="text-lg font-semibold">
                    {title}
                  </DialogTitle>
                  <DialogDescription className="text-gray-500 text-sm">
                    Submit button will activate when you make changes.
                  </DialogDescription>
                </div>

                {/* Scrollable Form Content */}
                <ScrollArea className="flex-1 px-6 py-4">
                  <div className="space-y-4">{children}</div>
                </ScrollArea>

                {/* Footer */}
                <div className="border-t px-6 py-4 flex justify-end">
                  <SubmitButton
                    text="Save Changes"
                    isLoading={loading}
                    disabled={!formState.isValid || !formState.isDirty}
                  />
                </div>
              </form>
            </FormProvider>
          )}
        </DialogContent>
      </Dialog>

      {/* Confirmation Dialog */}
      <Dialog open={showConfirmation} onOpenChange={setShowConfirmation}>
        <DialogContent className="sm:max-w-md p-6 rounded-xl space-y-6">
          <div className="text-center">
            <DialogTitle className="text-xl font-semibold text-gray-800">
              Unsaved Changes
            </DialogTitle>
            <p className="mt-2 text-gray-500 text-sm">
              You have made changes. Are you sure you want to discard them?
            </p>
          </div>

          <div className="flex justify-center gap-4">
            <Button
              variant="secondary"
              className="w-32"
              onClick={() => handleConfirmation(false)}
            >
              Keep Editing
            </Button>
            <Button
              variant="danger"
              className="w-32"
              onClick={() => handleConfirmation(true)}
            >
              Discard
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </>
  );
}
