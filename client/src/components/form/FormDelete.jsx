/* eslint-disable react/prop-types */
// src/components/form/FormDelete.jsx
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/dialog";
import { Trash } from "lucide-react";
import { Button } from "@/components/ui/button";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

const FormDelete = ({ title, description, onClick, loading = false }) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <button className="btn btn-danger">
          <Trash size={18} />
        </button>
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
                <Button variant="danger" className="w-32" onClick={onClick}>
                  Delete
                </Button>
              </DialogClose>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};

export { FormDelete };
