import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { FcGoogle } from "react-icons/fc";
import { Button } from "@/components/ui/button";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { FormInput } from "@/components/form/FormInput";
import { InputTextElement } from "@/components/input/InputTextElement";
import { registerSchema, sendOTPSchema, verifyOTPSchema } from "@/lib/schema";
import { sendOTPState, verifyOTPState, registerState } from "@/lib/constant";

const SignUp = () => {
  const { register, step, loading, resetStep } = useAuthStore();
  const [countdown, setCountdown] = useState(60);
  const [canResend, setCanResend] = useState(false);
  const [sentEmail, setSentEmail] = useState("");

  const getSchemaControl = () => {
    switch (step) {
      case 1:
        return sendOTPSchema;
      case 2:
        return verifyOTPSchema;
      case 3:
        return registerSchema;
      default:
        return [];
    }
  };

  const handleGoogleAuth = () => {
    window.location.href = `${import.meta.env.VITE_BASE_URL}/auth/google`;
  };

  const handleResendOTP = () => {
    if (!canResend || !sentEmail) return;
    register({ email: sentEmail });
    setCountdown(60);
    setCanResend(false);
  };

  useEffect(() => {
    resetStep();
  }, []);

  useEffect(() => {
    let timer;
    if (step === 2 && !canResend) {
      timer = setInterval(() => {
        setCountdown((prev) => {
          if (prev <= 1) {
            clearInterval(timer);
            setCanResend(true);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
    }
    return () => clearInterval(timer);
  }, [step, canResend]);

  return (
    <section className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="grid grid-cols-1 md:grid-cols-2 bg-white rounded-xl shadow-lg overflow-hidden max-w-4xl w-full">
        {/* Left Side */}
        <div className="hidden md:block bg-blue-600 p-8 text-white text-center">
          <h2 className="text-3xl font-bold mb-4">Welcome Back!</h2>
          <p className="text-sm">Stay Fit and Join us now</p>
          <img
            src="/signup-wallpaper.webp"
            alt="sign-up-illustration"
            className="mt-6 w-full h-auto"
          />
        </div>

        {/* Right Side */}
        <div className="p-8">
          <div className="mb-4">
            <WebLogo />
            <h3 className="text-center">Register</h3>
          </div>

          {/* Step Indicator */}
          <div className="flex justify-center mb-6 gap-2">
            {[1, 2, 3].map((s) => (
              <div
                key={s}
                className={`w-8 h-8 rounded-full flex items-center justify-center text-sm font-semibold ${
                  step === s
                    ? "bg-blue-600 text-white"
                    : "bg-gray-200 text-gray-600"
                }`}
              >
                {s}
              </div>
            ))}
          </div>

          <FormInput
            text={
              step === 1
                ? "Register with email"
                : step === 2
                ? "Submit OTP"
                : "Register"
            }
            action={(data) => {
              if (step === 1) setSentEmail(data.email);
              register(data);
            }}
            className="w-full py-3"
            isLoading={loading}
            state={
              step === 1
                ? sendOTPState
                : step === 2
                ? verifyOTPState
                : registerState
            }
            schema={getSchemaControl()}
          >
            {step === 1 && (
              <InputTextElement name="email" placeholder="Enter your email" />
            )}

            {step === 2 && (
              <>
                <div className="mb-2 text-sm text-gray-700">
                  We sent an OTP to:{" "}
                  <span className="font-medium">{sentEmail}</span>
                </div>
                <InputTextElement
                  name="otp"
                  isNumeric
                  maxLength={6}
                  label="OTP Code"
                  placeholder="Enter the OTP code"
                />
                <div className="text-sm text-right mt-2">
                  {canResend ? (
                    <button
                      onClick={handleResendOTP}
                      className="text-blue-600 hover:underline"
                    >
                      Resend OTP
                    </button>
                  ) : (
                    <span className="text-gray-500">
                      Resend in {countdown}s
                    </span>
                  )}
                </div>
              </>
            )}

            {step === 3 && (
              <div>
                <InputTextElement name="email" label="Email" disabled />
                <InputTextElement
                  name="password"
                  type="password"
                  label="Password"
                  placeholder="********"
                />
                <InputTextElement
                  name="fullname"
                  label="Fullname"
                  placeholder="Enter your fullname"
                />
              </div>
            )}
          </FormInput>

          {step === 1 && (
            <div>
              <div className="text-center py-2 text-sm">Or</div>
              <Button
                variant="outline"
                onClick={handleGoogleAuth}
                className="w-full"
              >
                <FcGoogle size={24} />
                Continue with Google
              </Button>
              <p className="text-sm text-center mt-6 text-gray-600">
                Already have an account?{" "}
                <Link
                  to="/signin"
                  className="text-blue-600 hover:underline font-medium"
                >
                  Login
                </Link>
              </p>
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

export default SignUp;
