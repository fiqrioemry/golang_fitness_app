import { OurClasses } from "@/components/home/OurClasses";
import { Membership } from "@/components/home/Membership";
import { HeroSection } from "@/components/home/HeroSection";
import { OurLocation } from "@/components/home/OurLocation";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";

const homeTitle = "Home -  high-intensity workouts,";
export default function Home() {
  useDocumentTitle(homeTitle);
  return (
    <main className="bg-background text-foreground">
      {/* hero section */}
      <HeroSection />

      {/* popular classes */}
      <OurClasses />

      {/* Membership*/}
      <Membership />

      {/* Visit Studio */}
      <OurLocation />
    </main>
  );
}
