import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/dialog";
import { Play, StopCircle, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

const FormToggle = ({
  title,
  description,
  onToggle,
  text = "start",
  toggle = true,
  loading = false,
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          className="w-full"
          variant={toggle ? "default" : "destructive"}
          type="button"
        >
          {toggle ? (
            <Play className="w-4 h-4" />
          ) : (
            <StopCircle className="w-4 h-4" />
          )}
          <span>{text}</span>
        </Button>
      </DialogTrigger>

      <DialogContent className="sm:max-w-md rounded-xl p-6 space-y-6">
        {loading ? (
          <SubmitLoading text="Deleting..." />
        ) : (
          <>
            <div className="text-center space-y-2">
              <DialogTitle className="text-2xl font-bold text-gray-800">
                {title}
              </DialogTitle>
              <DialogDescription className="text-gray-500">
                {description}
              </DialogDescription>
            </div>

            <div className="flex justify-center gap-4 pt-4">
              <DialogClose asChild>
                <Button variant="secondary" className="w-32">
                  Cancel
                </Button>
              </DialogClose>

              <DialogClose asChild>
                <Button
                  variant="destructive"
                  className="w-32"
                  onClick={onToggle}
                >
                  Yes, Execute
                </Button>
              </DialogClose>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};

export { FormToggle };
