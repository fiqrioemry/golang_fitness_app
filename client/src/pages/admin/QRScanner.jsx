// components/QRScanner.jsx
import React, { useEffect, useRef } from "react";
import { Html5QrcodeScanner } from "html5-qrcode";
import { useAttendanceMutation } from "@/hooks/useAttendance";

const QRScanner = () => {
  const scannerRef = useRef(null);
  const { validateQR } = useAttendanceMutation();

  useEffect(() => {
    const scanner = new Html5QrcodeScanner("qr-scanner", {
      fps: 10,
      qrbox: 250,
    });

    scanner.render(
      (decodedText) => {
        validateQR.mutate(decodedText);
        scanner.clear();
      },
      (err) => {
        console.warn("QR Scan Error:", err);
      }
    );

    return () => {
      scanner.clear().catch(console.error);
    };
  }, []);

  return (
    <div className="max-w-xl mx-auto py-12 px-4">
      <h2 className="text-2xl font-bold mb-6 text-center">
        Scan QR Code (Camera)
      </h2>
      <div id="qr-scanner" className="rounded-xl overflow-hidden border" />
    </div>
  );
};

export default QRScanner;
