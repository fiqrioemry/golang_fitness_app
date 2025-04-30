import React from "react";
import { Star } from "lucide-react";
import { useParams } from "react-router-dom";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { useClassDetailQuery } from "@/hooks/useClass";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const ClassDetail = () => {
  const { id } = useParams();
  const { data: cls, isLoading, isError, refetch } = useClassDetailQuery(id);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  console.log(cls);
  return (
    <section className="container mx-auto py-10 space-y-10">
      {/* Header Section */}
      <div className="flex flex-col lg:flex-row gap-8 items-start">
        <img
          src={cls.image}
          alt={cls.title}
          className="w-full lg:w-1/2 h-64 object-cover rounded-xl"
        />
        <div className="space-y-4 w-full">
          <h1 className="text-3xl font-bold">{cls.title}</h1>
          <p className="text-muted-foreground text-sm leading-relaxed">
            {cls.description}
          </p>
          <div className="flex flex-wrap gap-2">
            {cls.additional?.map((item, idx) => (
              <Badge key={idx} variant="outline">
                {item}
              </Badge>
            ))}
          </div>
          <div className="text-sm text-gray-500 space-y-1">
            <p>⏱ Duration: {cls.duration} minutes</p>
            <p>📍 Location: {cls.location}</p>
            <p>
              📌 Type: {cls.type} | Level: {cls.level}
            </p>
            <p>
              🏷 Category: {cls.category} - {cls.subcategory}
            </p>
          </div>
        </div>
      </div>

      {/* Gallery Section */}
      {cls.galleries?.length > 0 && (
        <div>
          <h2 className="text-xl font-semibold mb-4">Gallery</h2>
          <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
            {cls.galleries.map((url, idx) => (
              <img
                key={idx}
                src={url}
                alt={`Gallery ${idx + 1}`}
                className="w-full h-60   object-cover rounded-lg"
              />
            ))}
          </div>
        </div>
      )}

      {/* Review Section */}
      {cls.reviews?.length > 0 && (
        <div>
          <h2 className="text-xl font-semibold mb-4">Reviews</h2>
          <div className="space-y-4">
            {cls.reviews.map((review) => (
              <Card key={review.id}>
                <CardContent className="p-4">
                  <div className="flex justify-between items-center mb-1">
                    <p className="font-medium">{review.userName}</p>
                    <div className="flex items-center text-yellow-500 gap-1">
                      {[...Array(review.rating)].map((_, i) => (
                        <Star key={i} size={16} fill="currentColor" />
                      ))}
                    </div>
                  </div>
                  <p className="text-sm text-gray-600">{review.comment}</p>
                  <p className="text-xs text-gray-400 mt-1">
                    {new Date(review.createdAt).toLocaleDateString()}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      )}
    </section>
  );
};

export default ClassDetail;
