import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <main className="bg-white text-gray-900">
      {/* Hero Section */}
      <section
        className="relative w-full h-[800px] bg-cover bg-center"
        style={{ backgroundImage: "url('https://placehold.co/1920x800')" }}
      >
        <div className="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center">
          <div className="text-center px-4">
            <h1 className="text-5xl md:text-6xl font-bold text-white mb-4">
              Train Better. Live Stronger.
            </h1>
            <p className="text-xl text-white mb-6">
              Discover, schedule, and attend classes that energize your
              lifestyle — all in one beautifully simple platform.
            </p>
            <div className="flex flex-col md:flex-row gap-4 justify-center">
              <Button className="text-lg px-6 py-3"> Book a Class</Button>
              <Button variant="outline" className="text-lg px-6 py-3">
                See Our Packages
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Categories */}
      <section className="py-16 px-4 max-w-7xl mx-auto">
        <h2 className="text-3xl font-semibold text-center mb-4">
          Explore Our Classes
        </h2>
        <p className="text-center text-gray-600 mb-12 max-w-2xl mx-auto">
          From relaxing yoga sessions to high-intensity workouts, choose the
          class that matches your fitness goals and schedule.
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-8">
          {["Yoga", "HIIT", "Strength", "Zumba"].map((cat) => (
            <div
              key={cat}
              className="text-center border rounded-xl overflow-hidden shadow hover:shadow-lg transition"
            >
              <img
                src="https://placehold.co/400x250"
                alt={cat}
                className="w-full h-48 object-cover"
              />
              <div className="p-4">
                <h3 className="text-xl font-semibold">{cat}</h3>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Popular Classes */}
      <section className="py-16 px-4 bg-gray-100">
        <h2 className="text-3xl font-semibold text-center mb-12">
          Popular Classes
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto">
          {[1, 2, 3].map((_, i) => (
            <div
              key={i}
              className="bg-white rounded-xl overflow-hidden shadow hover:shadow-lg transition"
            >
              <img
                src="https://placehold.co/600x400"
                alt="class"
                className="w-full h-56 object-cover"
              />
              <div className="p-6">
                <h4 className="text-xl font-bold mb-2">Zumba Burn</h4>
                <p className="text-sm text-gray-600 mb-4">
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

      {/* Packages */}
      <section className="py-16 px-4 max-w-7xl mx-auto">
        <h2 className="text-3xl font-semibold text-center mb-12">
          Membership Packages
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-8">
          {["Basic", "Pro", "Unlimited"].map((pkg, i) => (
            <div
              key={pkg}
              className="bg-white border rounded-xl p-6 text-center shadow"
            >
              <h3 className="text-2xl font-bold mb-4">{pkg}</h3>
              <p className="text-gray-600 mb-4">
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

      {/* Visit Our Studio */}
      <section className="py-20 px-4 bg-white">
        <div className="max-w-6xl mx-auto space-y-10">
          <div className="text-center space-y-2">
            <h3 className="text-3xl font-bold text-gray-800">
              Visit Our Studio
            </h3>
            <p className="text-sm text-muted-foreground max-w-xl mx-auto">
              Experience the space where wellness begins. Drop by and feel the
              calm, the energy, and the inspiration.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8 items-start">
            {/* Map section */}
            <div className="w-full h-80 rounded-xl overflow-hidden shadow-lg ring-1 ring-gray-200">
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

            {/* Studio info */}
            <div className="p-6 space-y-4">
              <h4 className="text-xl font-semibold text-gray-800">
                Fitness Studio A
              </h4>
              <p className="text-sm text-gray-600">
                <span className="font-medium">Address:</span> 123 Fitness St,
                New York, NY
              </p>
              <p className="text-sm text-gray-600">
                <span className="font-medium">Hours:</span> Open daily from{" "}
                <strong>6AM – 10PM</strong>
              </p>
              <p className="text-sm text-gray-600">
                Come join our in-studio classes or simply drop in for a tour and
                meet our instructors in person!
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}
