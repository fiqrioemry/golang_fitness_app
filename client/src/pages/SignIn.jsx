import { Link } from "react-router-dom";
import { loginSchema } from "@/lib/schema";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { FormInput } from "@/components/form/FormInput";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { getLoginState } from "@/lib/constant";

const SignIn = () => {
  const { login, loading, rememberMe } = useAuthStore();

  return (
    <section className="min-h-screen flex items-center justify-center bg-background px-4">
      <div className="w-full max-w-4xl grid grid-cols-1 md:grid-cols-2 bg-card border border-border rounded-xl shadow-lg overflow-hidden">
        {/* Left Illustration (dengan bg-muted + aksen primary) */}

        <div className="hidden md:block relative h-[550px]">
          <img
            src="/login.png"
            alt="Wallpaper"
            className="absolute inset-0 w-full h-full object-cover"
          />
        </div>
        {/* <div className="hidden md:flex flex-col justify-center items-center bg-muted text-primary px-6 py-10">
          <h2 className="text-3xl font-bold mb-2">Welcome Back!</h2>
          <p className="text-sm text-muted-foreground text-center">
            Sign in and access your dashboard
          </p>
          <img
            src="/signin-wallpaper.webp"
            alt="sign-in"
            className="w-full h-auto mt-6 rounded"
          />
        </div> */}

        {/* Right Form */}
        <div className="px-6 py-10">
          <div className="mb-6 flex justify-center text-center">
            <WebLogo />
          </div>

          <FormInput
            action={login}
            text="Sign In"
            className="w-full"
            schema={loginSchema}
            isLoading={loading}
            state={getLoginState(rememberMe)}
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
              placeholder="********"
            />
            <SwitchElement name="rememberMe" label="Remember Me" />
          </FormInput>

          <p className="text-sm text-center mt-6 text-muted-foreground">
            Don't have an account?{" "}
            <Link
              to="/signup"
              className="text-primary font-medium hover:underline"
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
