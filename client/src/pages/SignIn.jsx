import { Link } from "react-router-dom";
import { loginSchema } from "@/lib/schema";
import { loginState } from "@/lib/constant";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { FormInput } from "@/components/form/FormInput";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";

const SignIn = () => {
  const { login, loading } = useAuthStore();

  return (
    <section className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="grid grid-cols-1 md:grid-cols-2 bg-white rounded-xl shadow-lg overflow-hidden max-w-4xl w-full">
        {/* Left Side (Illustration) */}
        <div className="hidden md:block bg-blue-600 p-8 text-white text-center">
          <h2 className="text-3xl font-bold mb-4">Welcome Back!</h2>
          <p className="text-sm">Sign in and access your dashboard</p>
          <img
            src="/signin-wallpaper.webp"
            alt="sign-in-illustration"
            className="mt-6 w-full h-auto"
          />
        </div>

        {/* Right Side (Form) */}
        <div className="p-8">
          <div className="mb-4">
            <WebLogo />
          </div>

          <FormInput
            action={login}
            text="Sign In"
            className="w-full"
            state={loginState}
            isLoading={loading}
            schema={loginSchema}
          >
            <InputTextElement
              name="email"
              label="Email"
              placeholder="Enter your email"
            />
            <InputTextElement
              name="password"
              label="Password"
              type="password"
              placeholder="*********"
            />
            <SwitchElement name="rememberMe" label="Remember Me" />
          </FormInput>

          <p className="text-sm text-center mt-6 text-gray-600">
            Don't have an account?{" "}
            <Link
              to="/signup"
              className="text-blue-600 hover:underline font-medium"
            >
              Sign up now
            </Link>
          </p>
        </div>
      </div>
    </section>
  );
};

export default SignIn;
