import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { api } from "../services/api";

export interface Stock {
  id: string;
  ticker: string;
  target_from: number;
  target_to: number;
  company: string;
  action: string;
  brokerage: string;
  rating_from: string;
  rating_to: string;
  time: string;
}

export interface StockIndicators {
  broker_confidence: number;
  price_target_growth: number;
  rating_impact: number;
}

export interface StockRecommendation {
  stock: Stock;
  score: number;
  indicators: StockIndicators;
  recommendation: string;
}

export const useStockStore = defineStore("stocks", () => {
  const stocks = ref<Stock[]>([]);
  const recommendedStocks = ref<StockRecommendation[]>([]);
  const selectedStock = ref<Stock | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  // Getters
  const sortedStocks = computed(() => {
    return [...stocks.value].sort((a, b) => {
      const changeA = ((a.target_to - a.target_from) / a.target_from) * 100;
      const changeB = ((b.target_to - b.target_from) / b.target_from) * 100;
      return changeB - changeA;
    });
  });

  const recommendedStocksList = computed(() => [...recommendedStocks.value]);

  // Actions
  async function fetchStocks() {
    loading.value = true;
    error.value = null;
    try {
      stocks.value = await api.get<Stock[]>("/api/stocks");
    } catch (e) {
      error.value = e instanceof Error ? e.message : "Error desconocido";
    } finally {
      loading.value = false;
    }
  }

  async function fetchRecommendedStocks() {
    loading.value = true;
    error.value = null;
    try {
      recommendedStocks.value = await api.get<StockRecommendation[]>(
        "/api/stocks/recommended"
      );
    } catch (e) {
      error.value = e instanceof Error ? e.message : "Error desconocido";
    } finally {
      loading.value = false;
    }
  }

  async function fetchStockDetail(symbol: string) {
    loading.value = true;
    error.value = null;
    try {
      selectedStock.value = await api.get<Stock>(`/api/stocks/${symbol}`);
    } catch (e) {
      error.value = e instanceof Error ? e.message : "Error desconocido";
    } finally {
      loading.value = false;
    }
  }

  async function syncStocks() {
    loading.value = true;
    error.value = null;
    try {
      await api.post("/api/stocks", {});
      // Reload stocks after syncing
      await fetchStocks();
      await fetchRecommendedStocks();
    } catch (e) {
      error.value = e instanceof Error ? e.message : "Error desconocido";
    } finally {
      loading.value = false;
    }
  }

  return {
    stocks,
    recommendedStocks,
    selectedStock,
    loading,
    error,
    sortedStocks,
    recommendedStocksList,
    fetchStocks,
    fetchRecommendedStocks,
    fetchStockDetail,
    syncStocks,
  };
});
