import React from "react";
import { Plus } from "lucide-react";
import DeleteClass from "./DeleteClass";
import UpdateClass from "./UpdateClass";
import { useNavigate } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { useClassesQuery } from "@/hooks/useClass";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const ClassesList = () => {
  const { data, isLoading, isError, refetch } = useClassesQuery();
  const classes = data?.classes || [];
  const navigate = useNavigate();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-8xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Classes Management</h2>
        <p className="text-muted-foreground text-sm">
          View, add, and manage training packages available for purchase by
          users.
        </p>
      </div>
      <div className="flex justify-end">
        <button
          onClick={() => navigate("/admin/classes/add")}
          className="flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90 transition"
        >
          <Plus className="w-4 h-4" />
          Tambah Kelas
        </button>
      </div>
      <div className="overflow-x-auto">
        <>
          {/* Desktop Table */}
          <div className="hidden md:block overflow-x-auto border rounded-xl shadow-sm">
            <table className="min-w-full bg-white text-sm">
              <thead className="bg-gray-100 text-gray-700 text-xs uppercase">
                <tr>
                  <th className="p-3 text-left">Thumbnail</th>
                  <th className="p-3 text-left">Judul</th>
                  <th className="p-3 text-left">Durasi</th>
                  <th className="p-3 text-left">Status</th>
                  <th className="p-3 text-left">Tanggal Dibuat</th>
                  <th className="p-3 text-left">Aksi</th>
                </tr>
              </thead>
              <tbody>
                {classes.map((classItem) => (
                  <tr key={classItem.id} className="border-t hover:bg-gray-50">
                    <td className="p-3">
                      <img
                        src={classItem.image}
                        alt={classItem.title}
                        className="w-16 h-16 object-cover rounded-md"
                      />
                    </td>
                    <td className="p-3 font-medium">{classItem.title}</td>
                    <td className="p-3">{classItem.duration} menit</td>
                    <td className="p-3">
                      <span
                        className={`px-2 py-1 rounded-full text-xs font-semibold ${
                          classItem.isActive
                            ? "bg-green-100 text-green-700"
                            : "bg-red-100 text-red-700"
                        }`}
                      >
                        {classItem.isActive ? "Aktif" : "Tidak Aktif"}
                      </span>
                    </td>
                    <td className="p-3">
                      {new Date(classItem.createdAt).toLocaleDateString()}
                    </td>
                    <td className="p-3">
                      <div className="flex gap-2">
                        <UpdateClass classes={classItem} />
                        <DeleteClass classes={classItem} />
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          {/* Mobile Card View */}
          <div className="md:hidden space-y-4">
            {classes.map((classItem) => (
              <div
                key={classItem.id}
                className="border rounded-lg p-4 shadow-sm space-y-2"
              >
                <div className="flex items-center gap-4">
                  <img
                    src={classItem.image}
                    alt={classItem.title}
                    className="w-16 h-16 object-cover rounded-md"
                  />
                  <div>
                    <h3 className="font-semibold text-base">
                      {classItem.title}
                    </h3>
                    <p className="text-xs text-muted-foreground">
                      {classItem.duration} menit
                    </p>
                  </div>
                </div>
                <div className="text-sm space-y-1">
                  <p className="text-muted-foreground">
                    Status:{" "}
                    <span
                      className={`px-2 py-0.5 rounded-full text-xs font-medium ${
                        classItem.isActive
                          ? "bg-green-100 text-green-700"
                          : "bg-red-100 text-red-700"
                      }`}
                    >
                      {classItem.isActive ? "Aktif" : "Tidak Aktif"}
                    </span>
                  </p>
                  <p className="text-muted-foreground">
                    Dibuat: {new Date(classItem.createdAt).toLocaleDateString()}
                  </p>
                </div>
                <div className="flex justify-end gap-2">
                  <UpdateClass classes={classItem} />
                  <DeleteClass classes={classItem} />
                </div>
              </div>
            ))}
          </div>
        </>
      </div>
    </section>
  );
};

export default ClassesList;
