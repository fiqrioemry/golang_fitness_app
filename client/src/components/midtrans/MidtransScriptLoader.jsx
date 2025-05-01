import { useEffect } from "react";

export const MidtransScriptLoader = () => {
  useEffect(() => {
    const existingScript = document.querySelector('script[src*="midtrans"]');
    if (existingScript) return;
    const script = document.createElement("script");
    script.src = import.meta.env.VITE_MIDTRANS_URL;
    script.setAttribute(
      "data-client-key",
      import.meta.env.VITE_MIDTRANS_PUBLIC_KEY
    );
    script.type = "text/javascript";
    script.async = true;
    document.body.appendChild(script);

    return () => {
      document.body.removeChild(script);
    };
  }, []);

  return null;
};
