import { create } from "zustand";

export const useBookingStore = create((set) => ({
  page: 1,
  limit: 10,
  sort: "date_desc",
  status: "upcoming",

  setPage: (val) => set({ page: val }),
  setStatus: (val) => set({ status: val }),
  setLimit: (val) => set({ limit: val }),
  setSort: (field) =>
    set((state) => ({
      sort: state.sort === `${field}_asc` ? `${field}_desc` : `${field}_asc`,
      page: 1,
    })),

  reset: () =>
    set({
      page: 1,
      limit: 10,
      role: "all",
      sort: "",
      status: "all",
      range: "daily",
    }),
}));
