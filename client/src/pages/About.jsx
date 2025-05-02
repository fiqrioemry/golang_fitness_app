import React from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useInstructorsQuery } from "@/hooks/useInstructor";
import { HeartPulse, CalendarClock, Users } from "lucide-react";

const About = () => {
  const {
    data: instructors = [],
    isLoading,
    isError,
    refetch,
  } = useInstructorsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="min-h-screen max-w-7xl mx-auto px-4 py-10 space-y-20">
      {/* Hero Section */}
      <div className="relative bg-gradient-to-r from-blue-500 to-indigo-600 rounded-xl text-white px-6 py-20 text-center shadow-lg">
        <h1 className="text-4xl font-bold mb-2">
          Empowering Your Wellness Journey
        </h1>
        <p className="text-blue-100 max-w-xl mx-auto text-sm">
          At Wellness Studio, we believe that fitness is not a goal—it's a
          lifestyle. Join our community and book classes tailored to your body
          and mind.
        </p>
      </div>

      {/* Our Mission */}
      <div className="text-center space-y-4">
        <h2 className="text-2xl font-bold">Our Mission</h2>
        <p className="text-muted-foreground max-w-xl mx-auto text-sm">
          To provide accessible, flexible, and enjoyable wellness classes that
          empower every individual to stay active, reduce stress, and live
          healthier.
        </p>
      </div>

      {/* Why Choose Us */}
      <div className="space-y-6">
        <h3 className="text-xl font-semibold text-center">
          Why Choose Wellness Studio?
        </h3>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 text-center">
          <div className="bg-white rounded-xl shadow p-6 space-y-2">
            <HeartPulse className="mx-auto text-blue-600 w-8 h-8" />
            <h4 className="font-semibold">Holistic Wellness</h4>
            <p className="text-sm text-gray-600">
              We focus on your physical, mental, and emotional well-being
              through diverse classes.
            </p>
          </div>
          <div className="bg-white rounded-xl shadow p-6 space-y-2">
            <CalendarClock className="mx-auto text-blue-600 w-8 h-8" />
            <h4 className="font-semibold">Flexible Scheduling</h4>
            <p className="text-sm text-gray-600">
              Book anytime, anywhere. Our platform is designed to fit your
              lifestyle.
            </p>
          </div>
          <div className="bg-white rounded-xl shadow p-6 space-y-2">
            <Users className="mx-auto text-blue-600 w-8 h-8" />
            <h4 className="font-semibold">Expert Instructors</h4>
            <p className="text-sm text-gray-600">
              All classes are guided by certified instructors who care about
              your progress.
            </p>
          </div>
        </div>
      </div>

      {/* Instructor Profiles */}
      <div className="space-y-6">
        <h3 className="text-xl font-semibold text-center">
          Meet Our Instructors
        </h3>
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-6">
          {instructors.map((ins) => (
            <div
              key={ins.id}
              className="bg-white shadow rounded-xl p-5 flex flex-col items-center text-center space-y-2"
            >
              <img
                src={ins.avatar}
                alt={ins.fullname}
                className="w-20 h-20 rounded-full border object-cover"
              />
              <h4 className="font-semibold text-sm">{ins.fullname}</h4>
              <p className="text-xs text-gray-500">{ins.specialties}</p>
            </div>
          ))}
        </div>
      </div>

      {/* CTA Section */}
      <div className="text-center bg-blue-50 rounded-xl p-10">
        <h3 className="text-xl font-semibold mb-2">Ready to feel your best?</h3>
        <p className="text-sm text-muted-foreground mb-4">
          Start your fitness journey with us today.
        </p>
        <Link to="/classes">
          <Button size="lg">Explore Classes</Button>
        </Link>
      </div>
    </section>
  );
};

export default About;
