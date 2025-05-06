import React from "react";
import { ArrowLeft, Star } from "lucide-react";
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

  return (
    <section className="section py-24 text-foreground">
      <div className="mb-6">
        <button
          onClick={() => history.back()}
          className="flex items-center text-sm text-muted-foreground hover:text-primary transition"
        >
          <ArrowLeft className="w-4 h-4 mr-1" />
          Back to Classes
        </button>
      </div>
      {/* Header Section */}
      <div className="flex flex-col lg:flex-row gap-8 items-start">
        <img
          src={cls.image}
          alt={cls.title}
          className="w-full lg:w-1/2 h-64 object-cover rounded-2xl border"
        />
        <div className="space-y-4 w-full">
          <h2 className="text-3xl font-bold text-foreground">{cls.title}</h2>
          <p className="text-subtitle leading-relaxed">{cls.description}</p>

          <div className="flex flex-wrap gap-2">
            {cls.additional?.map((item, idx) => (
              <Badge key={idx} variant="outline">
                {item}
              </Badge>
            ))}
          </div>

          <div className="text-sm text-muted-foreground space-y-1">
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
          <h2 className="text-xl font-semibold text-foreground mb-4">
            Gallery
          </h2>
          <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
            {cls.galleries.map((url, idx) => (
              <img
                key={idx}
                src={url}
                alt={`Gallery ${idx + 1}`}
                className="w-full h-60 object-cover rounded-xl border"
              />
            ))}
          </div>
        </div>
      )}

      {/* Review Section */}
      {cls.reviews?.length > 0 && (
        <div>
          <h2 className="text-xl font-semibold text-foreground mb-4">
            Reviews
          </h2>
          <div className="space-y-4">
            {cls.reviews.map((review) => (
              <Card key={review.id} className="card justify-start">
                <CardContent className="p-4 flex items-start">
                  <div className="flex justify-between items-center mb-1">
                    <p className="font-medium text-foreground">
                      {review.userName}
                    </p>
                    <div className="flex items-center text-yellow-500 gap-1">
                      {[...Array(review.rating)].map((_, i) => (
                        <Star key={i} size={16} fill="currentColor" />
                      ))}
                    </div>
                  </div>
                  <p className="text-sm text-subtitle">{review.comment}</p>
                  <p className="text-xs text-muted-foreground mt-1">
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
