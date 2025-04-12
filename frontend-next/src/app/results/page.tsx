'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Button } from '@/components/ui/button'
import { ResultCard } from '@/components/results/ResultCard'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { ErrorDisplay } from '@/components/ui/ErrorDisplay'

type Result = {
  id: number
  quiz_id: number
  score: number
  total_questions: number
  correct_answers: number
  time_taken: number
  is_passed: boolean
  passing_score: number
  created_at: string
}

export default function ResultsPage() {
  const [results, setResults] = useState<Result[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const router = useRouter()

  const fetchResults = async () => {
    try {
      setLoading(true)
      setError(null)
      const res = await fetch('/api/results')
      if (!res.ok) {
        throw new Error('Failed to fetch results')
      }
      const data = await res.json()
      setResults(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load results')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchResults()
  }, [])

  if (loading) {
    return <LoadingSpinner className="min-h-screen" text="Loading your results..." />
  }

  if (error) {
    return (
      <ErrorDisplay 
        error={error}
        onRetry={fetchResults}
        className="min-h-screen"
      />
    )
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold bg-gradient-to-r from-blue-500 to-purple-500 bg-clip-text text-transparent">
          Your Quiz Results
        </h1>
        <Button 
          onClick={() => router.push('/')}
          className="bg-blue-500 hover:bg-blue-600 text-white"
        >
          Take Another Quiz
        </Button>
      </div>
      
      {results.length === 0 ? (
        <div className="text-center py-12 bg-gray-50 rounded-lg">
          <div className="text-6xl mb-4">📝</div>
          <p className="text-gray-600 mb-4">No results found yet</p>
          <Button 
            onClick={() => router.push('/')}
            className="bg-blue-500 hover:bg-blue-600 text-white"
          >
            Take Your First Quiz
          </Button>
        </div>
      ) : (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {results.map((result) => (
            <ResultCard key={result.id} result={result} />
          ))}
        </div>
      )}
    </div>
  )
}
