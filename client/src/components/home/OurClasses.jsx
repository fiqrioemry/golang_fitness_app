import React from "react";
import { Button } from "@/components/ui/button";

const OurClasses = () => {
  return (
    <section className="py-20 px-4 bg-muted">
      <h2 className="text-4xl font-bold text-center mb-4 font-heading">
        Popular Classes
      </h2>
      <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
        From relaxing yoga sessions to high-intensity workouts, choose the class
        that suits your fitness goals.
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
  );
};

export { OurClasses };
