import {
  Sheet,
  SheetTrigger,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetDescription,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { ScrollArea } from "@/components/ui/scroll-area";
import { useEffect, useMemo, useCallback, useState } from "react";
import { SubmitButton } from "@/components/form/SubmitButton";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

export function FormSheet({
  title,
  state,
  schema,
  action,
  children,
  open,
  setOpen,
  resourceId = null,
  loading = false,
  shouldReset = true,
}) {
  const [showConfirmation, setShowConfirmation] = useState(false);

  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
    mode: "onChange",
  });

  const { formState, reset, handleSubmit } = methods;
  const isFormDirty = useMemo(() => formState.isDirty, [formState.isDirty]);

  const resetAndCloseSheet = useCallback(() => {
    reset();
    setOpen(false);
  }, [reset, setOpen]);

  const handleCancel = useCallback(() => {
    if (isFormDirty) setShowConfirmation(true);
    else resetAndCloseSheet();
  }, [isFormDirty, resetAndCloseSheet]);

  const handleConfirmation = useCallback(
    (confirmed) => {
      if (confirmed) resetAndCloseSheet();
      setShowConfirmation(false);
    },
    [resetAndCloseSheet]
  );

  useEffect(() => {
    if (state) reset(state);
  }, [state, reset]);

  const handleSave = useCallback(
    async (data) => {
      if (resourceId) {
        await action({ id: resourceId, data });
      } else {
        await action(data);
      }

      if (formState.isValid && shouldReset) reset();
      setOpen(false);
    },
    [action, formState.isValid, reset, shouldReset, resourceId, setOpen]
  );

  return (
    <>
      {showConfirmation && (
        <div className="fixed inset-0 bg-black/30 flex items-center justify-center">
          <div className="bg-white rounded-lg p-6 space-y-4 max-w-sm w-full">
            <div className="text-center">
              <h2 className="text-lg font-semibold">Unsaved Changes</h2>
              <p className="text-sm text-gray-500">
                You have made changes. Are you sure you want to discard them?
              </p>
            </div>
            <div className="flex justify-center gap-4">
              <Button
                variant="secondary"
                onClick={() => handleConfirmation(false)}
              >
                Keep Editing
              </Button>
              <Button
                variant="destructive"
                onClick={() => handleConfirmation(true)}
              >
                Discard
              </Button>
            </div>
          </div>
        </div>
      )}
      <Sheet
        open={open}
        onOpenChange={(val) => (!val ? handleCancel() : setOpen(val))}
      >
        <SheetContent className="w-full sm:max-w-lg flex flex-col p-0">
          {loading ? (
            <SubmitLoading />
          ) : (
            <FormProvider {...methods}>
              <form
                onSubmit={handleSubmit(handleSave)}
                className="flex flex-col h-full"
              >
                <div className="border-b px-6 py-4">
                  <SheetHeader>
                    <SheetTitle className="text-lg font-semibold">
                      {title}
                    </SheetTitle>
                    <SheetDescription className="text-sm text-muted-foreground">
                      Submit button will activate when you make changes.
                    </SheetDescription>
                  </SheetHeader>
                </div>

                <ScrollArea className="flex-1 px-6 py-4">
                  <div className="space-y-4">{children}</div>
                </ScrollArea>

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
        </SheetContent>
      </Sheet>
    </>
  );
}
