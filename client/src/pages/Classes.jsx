import React from "react";
import { Link } from "react-router-dom";
import { useClassesQuery } from "../hooks/useClass";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

const Classes = () => {
  const { data: response, isLoading, isError, refetch } = useClassesQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const { classes = [] } = response;

  return (
    <section className="container mx-auto py-10">
      <h1 className="text-3xl font-bold mb-6 text-center">
        Explore Our Classes
      </h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        {classes.map((cls) => (
          <Link to={`/classes/${cls.id}`} key={cls.id}>
            <Card className="shadow-md hover:shadow-lg transition-shadow cursor-pointer h-full">
              <img
                src={cls.image}
                alt={cls.title}
                className="w-full h-48 object-cover rounded-t-xl"
              />
              <CardContent className="p-4 space-y-2">
                <h3 className="text-xl font-semibold line-clamp-2">
                  {cls.title}
                </h3>
                <p className="text-sm text-muted-foreground line-clamp-3">
                  {cls.description}
                </p>
                <div className="flex flex-wrap gap-2 py-2">
                  {cls.additional?.map((tag, idx) => (
                    <Badge key={idx} variant="outline">
                      {tag}
                    </Badge>
                  ))}
                </div>
                <div className="text-xs text-gray-500 pt-2">
                  Duration: {cls.duration} mins
                </div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </section>
  );
};

export default Classes;
