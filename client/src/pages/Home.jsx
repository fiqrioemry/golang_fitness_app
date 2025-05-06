import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { useEffect, useState } from "react";

export default function Home() {
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
    <main className="bg-background text-foreground">
      <section
        className="relative w-full h-[800px] bg-cover bg-center transition-all duration-1000 ease-in-out"
        style={{ backgroundImage: `url('${heroImages[currentImage]}')` }}
      >
        <div className="absolute inset-0 bg-black/50 flex items-center justify-center">
          <div className="text-center px-4 max-w-3xl">
            <h1 className="text-4xl md:text-6xl font-bold text-primary/90 mb-4 font-heading">
              Train Better. Live Stronger.
            </h1>
            <p className="text-lg text-muted-foreground mb-6">
              Discover, schedule, and attend classes that energize your
              lifestyle — all in one platform.
            </p>
            <div className="flex flex-col md:flex-row gap-4 justify-center">
              <Button size="lg" className="w-60 h-12 rounded-full">
                Book a Class
              </Button>
              <Button
                size="lg"
                variant="outline"
                className="w-60 h-12 rounded-full"
              >
                See Our Packages
              </Button>
            </div>
          </div>
        </div>
      </section>

      <section className="py-24 px-4 max-w-7xl mx-auto">
        <h2 className="text-4xl font-bold text-center mb-4 font-heading">
          Explore Our Program
        </h2>
        <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
          From relaxing yoga sessions to high-intensity workouts, choose the
          class that suits your fitness goals.
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

      {/* Popular Classes - desain lama */}
      <section className="py-20 px-4 bg-muted">
        <h2 className="text-4xl font-bold text-center mb-4 font-heading">
          Popular Classes
        </h2>
        <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
          From relaxing yoga sessions to high-intensity workouts, choose the
          class that suits your fitness goals.
        </p>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto">
          {[1, 2, 3].map((_, i) => (
            <div
              key={i}
              className="bg-card text-foreground border border-border rounded-xl shadow hover:shadow-xl transition"
            >
              <img
                src="/class-image1.webp"
                alt="class"
                className="w-full h-56 object-cover"
              />
              <div className="p-6">
                <h4 className="text-lg font-bold mb-2">Zumba Burn</h4>
                <p className="text-sm text-muted-foreground mb-4">
                  Today at 6 PM · Central Studio
                </p>
                <div className="flex justify-between items-center">
                  <span className="font-bold text-lg">Rp 75.000</span>
                  <Button size="sm">Book Now</Button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Membership Packages - desain lama */}
      <section className="py-20 px-4 max-w-7xl mx-auto">
        <h2 className="text-4xl font-bold text-center mb-4 font-heading">
          Membership Packages
        </h2>
        <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
          From relaxing yoga sessions to high-intensity workouts, choose the
          class that suits your fitness goals.
        </p>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-6">
          {["Basic", "Pro", "Unlimited"].map((pkg) => (
            <div
              key={pkg}
              className="bg-card text-foreground border border-border rounded-xl p-6 text-center shadow"
            >
              <h3 className="text-2xl font-bold mb-4">{pkg}</h3>
              <p className="text-sm text-muted-foreground mb-4">
                {pkg === "Basic"
                  ? "4 classes/month"
                  : pkg === "Pro"
                  ? "12 classes/month"
                  : "Unlimited classes"}
              </p>
              <span className="text-3xl font-bold block mb-4">
                {pkg === "Basic"
                  ? "Rp 250K"
                  : pkg === "Pro"
                  ? "Rp 600K"
                  : "Rp 950K"}
              </span>
              <Button className="w-full">Buy Package</Button>
            </div>
          ))}
        </div>
      </section>

      {/* Visit Studio */}
      <section className="py-20 px-4 bg-muted border-t border-border">
        <div className="max-w-6xl mx-auto space-y-10">
          <div className="text-center space-y-2">
            <h2 className="text-4xl font-bold text-center mb-4 font-heading">
              Visit Our Studio
            </h2>
            <p className="text-sm text-muted-foreground max-w-xl mx-auto">
              Experience our studio space — drop in, feel the energy, and meet
              our instructors in person.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="w-full h-80 rounded-xl overflow-hidden shadow ring-1 ring-border">
              <iframe
                title="Studio Map"
                width="100%"
                height="100%"
                frameBorder="0"
                style={{ border: 0 }}
                loading="lazy"
                allowFullScreen
                src="https://www.google.com/maps?q=40.712776,-74.005974&z=15&output=embed"
              ></iframe>
            </div>
            <div className="p-6 space-y-4">
              <h4 className="text-lg font-semibold">Fitness Studio A</h4>
              <p className="text-sm text-muted-foreground">
                <span className="font-medium">Address:</span> 123 Fitness St,
                New York, NY
              </p>
              <p className="text-sm text-muted-foreground">
                <span className="font-medium">Hours:</span> 6AM – 10PM daily
              </p>
              <p className="text-sm text-muted-foreground">
                Join a class, get a tour, or just stop by to see what we’re all
                about!
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}
