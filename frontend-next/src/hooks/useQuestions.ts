import { useEffect, useState } from "react"
import { apiRequest, API_ENDPOINTS } from "@/lib/api"

// This type defines what a quiz question looks like
// Each question has a unique ID, the question text, and a category
type Question = {
  id: number
  question: string
  category: string
}

// This hook helps us get quiz questions from the server
// It handles loading states, errors, and storing the questions
export const useQuestions = () => {
  // Store the list of questions
  const [questions, setQuestions] = useState<Question[]>([])
  // Track if we're loading questions
  const [loading, setLoading] = useState(true)
  // Store any error that happens
  const [error, setError] = useState<string | null>(null)

  // This runs when the component using this hook first loads
  useEffect(() => {
    // Function to get questions from the server
    const fetchQuestions = async () => {
      try {
        setLoading(true)
        setError(null)

        // Try to get questions from our API
        const { data, error } = await apiRequest<{ questions: Question[] }>(
          API_ENDPOINTS.questions
        )

        if (error) {
          throw new Error(error)
        }

        if (!data?.questions || !Array.isArray(data.questions)) {
          throw new Error("Invalid response format")
        }

        // Make sure each question has the right data
        const validQuestions = data.questions.filter((q) => 
          typeof q.id === 'number' && 
          typeof q.question === 'string' && 
          typeof q.category === 'string'
        )

        // Make sure we got at least one valid question
        if (validQuestions.length === 0) {
          throw new Error("No valid questions found in response")
        }

        // Save the questions we got
        setQuestions(validQuestions)
      } catch (err) {
        // If anything goes wrong, log it and save the error
        console.error("Error fetching questions:", err)
        setError(err instanceof Error ? err.message : "Failed to load questions")
      } finally {
        // Whether it worked or not, we're done loading
        setLoading(false)
      }
    }

    // Start getting the questions
    fetchQuestions()
  }, []) // Empty array means this only runs once when component mounts

  // Give back everything needed to use the questions
  return { questions, loading, error }
}
