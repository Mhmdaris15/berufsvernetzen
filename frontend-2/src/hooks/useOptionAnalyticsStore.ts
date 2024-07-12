import { create } from "zustand";

type OptionAnalyticsStore = {
  working: boolean;
  study: boolean;
  entrepreneurship: boolean;
  unemployment: boolean;
  all: boolean; 
  setWorking: () => void;
  setStudy: () => void;
  setEntrepreneurship: () => void;
  setUnemployment: () => void;
  setAll: () => void; // new
  setOneOption: (option: string) => void; // new
};

export const useOptionAnalyticsStore = create<OptionAnalyticsStore>(
  (set) => ({
    working: false,
    study: false,
    entrepreneurship: false,
    unemployment: false,
    all: true, // new
    setWorking: () => set((state) => ({ working: !state.working })),
    setStudy: () => set((state) => ({ study: !state.study })),
    setEntrepreneurship: () => set((state) => ({ entrepreneurship: !state.entrepreneurship })),
    setUnemployment: () => set((state) => ({ unemployment: !state.unemployment })),
    setAll: () => set((state) => ({ all: !state.all })), // new
    // make one option true, and others false
    setOneOption: (option) => {
      set((state) => ({
        working: option === "working",
        study: option === "study",
        entrepreneurship: option === "entrepreneurship",
        unemployment: option === "unemployment",
        all: option === "all",
      }));
    },
  })
);