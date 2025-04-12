import { apiRequest, API_ENDPOINTS } from '../api'

// Mock fetch
global.fetch = jest.fn()

describe('API Client', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  describe('apiRequest', () => {
    it('should make a successful API request', async () => {
      const mockData = { message: 'success' }
      const mockResponse = {
        ok: true,
        json: () => Promise.resolve(mockData),
      }
      ;(global.fetch as jest.Mock).mockResolvedValue(mockResponse)

      const result = await apiRequest('/test')

      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/test'),
        expect.any(Object)
      )
      expect(result.data).toEqual(mockData)
      expect(result.error).toBeUndefined()
    })

    it('should handle API errors', async () => {
      const mockResponse = {
        ok: false,
        status: 404,
      }
      ;(global.fetch as jest.Mock).mockResolvedValue(mockResponse)

      const result = await apiRequest('/test')

      expect(result.data).toBeUndefined()
      expect(result.error).toBe('HTTP error! status: 404')
    })

    it('should handle network errors', async () => {
      ;(global.fetch as jest.Mock).mockRejectedValue(
        new Error('Network error')
      )

      const result = await apiRequest('/test')

      expect(result.data).toBeUndefined()
      expect(result.error).toBe('Network error')
    })
  })

  describe('API_ENDPOINTS', () => {
    it('should have all required endpoints', () => {
      expect(API_ENDPOINTS).toEqual({
        questions: '/questions',
        results: '/results',
        users: '/users',
        auth: '/auth',
      })
    })
  })
}) 