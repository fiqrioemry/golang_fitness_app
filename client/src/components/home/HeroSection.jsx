import React from "react";
import { useEffect, useState } from "react";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";

const HeroSection = () => {
  const heroImages = [
    "/hero-image1.webp",
    "/hero-image2.webp",
    "/hero-image3.webp",
    "/hero-image4.webp",
  ];

  const [currentImage, setCurrentImage] = useState(0);
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentImage((prev) => (prev + 1) % heroImages.length);
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <section className="py-24 px-4 max-w-7xl mx-auto">
      <h2 className="text-4xl font-bold text-center mb-4 font-heading">
        Explore Our Program
      </h2>
      <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
        From relaxing yoga sessions to high-intensity workouts, choose the class
        that suits your fitness goals.
      </p>

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-8">
        {[
          { name: "Yoga", image: "/class-image1.webp" },
          { name: "HIIT", image: "/class-image2.webp" },
          { name: "Strength", image: "/class-image3.webp" },
          { name: "Zumba", image: "/class-image4.webp" },
        ].map((cls) => (
          <Card
            key={cls.name}
            className="group overflow-hidden rounded-2xl shadow hover:shadow-xl transition"
          >
            <div className="relative w-full h-48 overflow-hidden">
              <img
                src={cls.image}
                alt={cls.name}
                className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div className="absolute inset-0 bg-black/30 group-hover:bg-black/40 transition" />
            </div>
            <CardHeader className="bg-background text-center py-4">
              <CardTitle className="text-lg font-semibold text-foreground">
                {cls.name}
              </CardTitle>
            </CardHeader>
          </Card>
        ))}
      </div>
    </section>
  );
};

export { HeroSection };
