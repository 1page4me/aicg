/**
 * Base URL for API requests
 * @default 'http://localhost:8080/api'
 */
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'

/**
 * Interface for API response
 * @template T - Type of the data returned by the API
 */
export interface ApiResponse<T> {
  data?: T
  error?: string
}

/**
 * Makes an API request to the backend
 * @template T - Expected response type
 * @param endpoint - API endpoint to call
 * @param options - Fetch options
 * @returns Promise with API response
 */
export async function apiRequest<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  try {
    const response = await fetch(`${API_URL}${endpoint}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()
    return { data }
  } catch (error) {
    return {
      error: error instanceof Error ? error.message : 'An unknown error occurred',
    }
  }
}

/**
 * API endpoint constants
 */
export const API_ENDPOINTS = {
  questions: '/questions',
  results: '/results',
  users: '/users',
  auth: '/auth',
} as const 