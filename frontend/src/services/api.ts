const API_URL = import.meta.env.VITE_API_URL;

export class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = "ApiError";
  }
}

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ message: "Error desconocido" }));
    throw new ApiError(
      response.status,
      error.message || `Error HTTP: ${response.status}`
    );
  }
  return response.json();
}

const defaultConfig = {
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
};

export const api = {
  async get<T>(endpoint: string): Promise<T> {
    const response = await fetch(`${API_URL}${endpoint}`, {
      ...defaultConfig,
      method: "GET",
    });
    return handleResponse<T>(response);
  },

  async post<T>(endpoint: string, data: any): Promise<T> {
    const response = await fetch(`${API_URL}${endpoint}`, {
      ...defaultConfig,
      method: "POST",
      body: JSON.stringify(data || {}),
    });
    return handleResponse<T>(response);
  },
};
