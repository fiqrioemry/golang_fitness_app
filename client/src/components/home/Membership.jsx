import { Button } from "@/components/ui/Button";

export const Membership = () => {
  return (
    <section className="py-20 px-4 max-w-7xl mx-auto">
      <h2 className="text-4xl font-bold text-center mb-4 font-heading">
        Membership Packages
      </h2>
      <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
        From relaxing yoga sessions to high-intensity workouts, choose the class
        that suits your fitness goals.
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
  );
};
