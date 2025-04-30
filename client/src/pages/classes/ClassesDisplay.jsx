import React from "react";
import DeleteClass from "./DeleteClass";
import UpdateClass from "./UpdateClass";
import { useNavigate } from "react-router-dom";
import { useClassesQuery } from "@/hooks/useClass";
import { Loading } from "@/components/ui/Loading";
import { Pencil, Trash2, Plus } from "lucide-react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const ClassesDisplay = () => {
  const { data, isLoading, isError, refetch } = useClassesQuery();
  const classes = data?.classes || [];
  const navigate = useNavigate();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="p-6">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold">Manajemen Kelas</h2>
        <button
          onClick={() => navigate("/classes/add")}
          className="flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90 transition"
        >
          <Plus className="w-4 h-4" />
          Tambah Kelas
        </button>
      </div>

      <div className="overflow-x-auto">
        <table className="min-w-full bg-white rounded-md shadow-sm text-sm">
          <thead className="bg-gray-100 text-gray-700">
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
              <tr key={classItem.id} className="border-t  hover:bg-gray-50">
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
    </section>
  );
};

export default ClassesDisplay;
