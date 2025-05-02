import React from "react";
import { Link } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { useSearchParams } from "react-router-dom";
import { useClassesQuery } from "@/hooks/useClass";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import FilterSelection from "@/components/input/FilterSelection";

const Classes = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const filters = {
    typeId: searchParams.get("typeId"),
    levelId: searchParams.get("levelId"),
    categoryId: searchParams.get("categoryId"),
    locationId: searchParams.get("locationId"),
    subcategoryId: searchParams.get("subcategoryId"),
  };
  const {
    data: response,
    isLoading,
    isError,
    refetch,
  } = useClassesQuery(filters);

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const { classes = [] } = response;

  return (
    <section className="min-h-screen px-4 py-10 max-w-7xl mx-auto space-y-8">
      {/* Heading */}
      <div className="bg-gradient-to-r from-sky-500 to-indigo-600 text-white rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <div className="max-w-4xl mx-auto text-center space-y-2">
          <h3 className="text-3xl font-bold flex items-center justify-center gap-2">
            Explore Fitness Classes
          </h3>
          <p className="text-sm text-blue-100">
            Discover personalized sessions tailored for your needs, from
            beginner to advanced levels.
          </p>
        </div>
      </div>

      {/* Filter Bar */}
      <div className="sticky top-4 z-10 bg-white border shadow-sm rounded-xl p-4 mb-8">
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
          <FilterSelection
            paramKey="locationId"
            label="Location"
            data="location"
          />
          <FilterSelection paramKey="typeId" label="Type" data="type" />
          <FilterSelection
            paramKey="categoryId"
            label="Category"
            data="category"
          />
          <FilterSelection
            paramKey="subcategoryId"
            label="Subcategory"
            data="subcategory"
          />
          <FilterSelection paramKey="levelId" label="Level" data="level" />
        </div>
      </div>

      {classes.length === 0 ? (
        <div className="flex flex-col items-center justify-center text-center py-4 col-span-full">
          <img
            src="/no-classes.webp"
            alt="No class found"
            className="w-72 mb-6 opacity-70"
          />
          <h3 className="text-xl font-semibold mb-2">No classes found</h3>
          <p className="text-gray-500 mb-4 max-w-md text-sm">
            We couldn't find any classes based on your current filters. Try
            adjusting your selections or reset all filters.
          </p>
          <button
            onClick={() => setSearchParams({})}
            className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition"
          >
            Reset Filters
          </button>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {classes.map((cls) => (
            <Link to={`/classes/${cls.id}`} key={cls.id}>
              <div className="relative bg-white rounded-xl shadow-md hover:shadow-xl transition-transform hover:-translate-y-1 duration-300 group h-full flex flex-col">
                <div className="relative h-48 w-full overflow-hidden">
                  <img
                    src={cls.image}
                    alt={cls.title}
                    className={`w-full h-full object-cover ${
                      !cls.isActive
                        ? "grayscale brightness-75"
                        : "group-hover:scale-105 transition-all"
                    }`}
                  />
                  {!cls.isActive && (
                    <div className="absolute inset-0 bg-black/40 backdrop-blur-sm flex items-center justify-center z-10">
                      <span className="bg-red-600 text-white text-xs font-bold uppercase px-3 py-1 rounded-full">
                        Registration Closed
                      </span>
                    </div>
                  )}
                </div>

                <div className="flex flex-col justify-between flex-1 p-4">
                  <div>
                    <h3 className="text-lg font-semibold text-gray-800 mb-1 line-clamp-2">
                      {cls.title}
                    </h3>
                    <p className="text-sm text-muted-foreground mb-2 line-clamp-2">
                      {cls.description}
                    </p>
                  </div>

                  <div className="text-xs text-gray-500 mt-auto pt-2">
                    Duration: {cls.duration} mins
                  </div>
                </div>
              </div>
            </Link>
          ))}
        </div>
      )}
    </section>
  );
};

export default Classes;
