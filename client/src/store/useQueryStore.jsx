import { create } from "zustand";

export const useQueryStore = create((set) => ({
  q: "",
  page: 1,
  limit: 10,
  role: "all",
  sort: "",
  status: "all",

  setPage: (val) => set({ page: val }),
  setStatus: (val) => set({ status: val }),
  setLimit: (val) => set({ limit: val }),
  setRole: (val) => set({ role: val, page: 1 }),
  setQ: (val) => set({ q: val, page: 1 }),
  setSort: (field) =>
    set((state) => ({
      sort: state.sort === `${field}_asc` ? `${field}_desc` : `${field}_asc`,
      page: 1,
    })),

  reset: () =>
    set({
      q: "",
      page: 1,
      limit: 10,
      role: "all",
      sort: "",
      status: "all",
    }),
}));
