import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogFooter,
  DialogContent,
  DialogDescription,
} from "@/components/ui/dialog";

import { Button } from "./button";
import { useState, useEffect } from "react";
import { AlertTriangle } from "lucide-react";

const ErrorDialog = ({ open = true, onRetry }) => {
  const [visible, setVisible] = useState(open);

  useEffect(() => {
    if (!visible && onRetry) {
      setTimeout(() => {
        onRetry();
      }, 500); // auto refetch setelah close
    }
  }, [visible, onRetry]);

  return (
    <Dialog open={visible} onOpenChange={setVisible}>
      <DialogContent className="max-w-md sm:rounded-2xl sm:p-6 shadow-lg">
        <DialogHeader className="flex flex-col items-center text-center space-y-2">
          <div className="flex items-center justify-center w-16 h-16 rounded-full bg-red-100 text-red-600">
            <AlertTriangle className="w-8 h-8" />
          </div>
          <DialogTitle className="text-xl font-semibold text-red-600">
            Gagal Memuat Data
          </DialogTitle>
          <DialogDescription className="text-muted-foreground text-sm max-w-xs">
            Terjadi kesalahan saat mengambil data produk. Coba periksa koneksi
            internet kamu, lalu klik tombol di bawah untuk memuat ulang halaman.
          </DialogDescription>
        </DialogHeader>

        <DialogFooter className="mt-4 w-full">
          <Button className="w-full" onClick={() => setVisible(false)}>
            Coba Lagi
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default ErrorDialog;
