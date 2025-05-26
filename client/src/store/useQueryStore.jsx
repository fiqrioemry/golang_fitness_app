import { create } from "zustand";

export const useQueryStore = create((set) => ({
  page: 1,
  limit: 10,
  role: "all",
  sort: "",
  status: "all",
  range: "daily",

  setPage: (val) => set({ page: val }),
  setRange: (val) => set({ range: val }),
  setStatus: (val) => set({ status: val }),
  setLimit: (val) => set({ limit: val }),
  setRole: (val) => set({ role: val, page: 1 }),
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
