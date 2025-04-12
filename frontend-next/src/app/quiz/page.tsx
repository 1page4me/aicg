"use client"

import React, { useState } from "react"
import { useQuestions } from "@/hooks/useQuestions"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { motion } from "framer-motion"

// This is the main quiz page component
export default function QuizPage() {
  // Get questions and loading/error states from our custom hook
  const { questions, loading, error } = useQuestions()
  // Keep track of which question we're showing
  const [currentIndex, setCurrentIndex] = useState(0)

  // Show loading message while questions are being fetched
  if (loading) return <p className="p-4 text-muted">Loading...</p>
  // Show error message if something went wrong
  if (error) return <p className="p-4 text-red-500">{error}</p>

  // Get the current question to display
  const currentQuestion = questions[currentIndex]
  // Calculate how far along in the quiz we are (0-100%)
  const progress = Math.round((currentIndex / questions.length) * 100)

  // Handle when user answers a question
  const handleAnswer = (response: string) => {
    console.log("Answered:", currentQuestion.question, "->", response)
    // If there are more questions, go to next one
    if (currentIndex < questions.length - 1) {
      setCurrentIndex(currentIndex + 1)
    } else {
      // If this was the last question, show completion message
      alert("🎉 Quiz complete!")
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-sky-100 via-white to-indigo-100 p-6">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold text-center mb-4 text-indigo-600">Discover Your Strengths 🚀</h1>
        <Progress value={progress} className="mb-6 h-3 rounded-full" />

        <motion.div
          key={currentQuestion.id}
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -30 }}
          transition={{ duration: 0.4 }}
        >
          <Card className="bg-white/90 shadow-xl border-2 border-indigo-200">
            <CardContent className="p-8 space-y-6">
              <h2 className="text-xl font-semibold text-gray-800">
                Q{currentIndex + 1}. {currentQuestion.question}
              </h2>
              <div className="flex flex-col gap-4 sm:flex-row sm:justify-around">
                <Button
                  variant="outline"
                  className="w-full sm:w-auto text-green-700 border-green-400 hover:bg-green-100"
                  onClick={() => handleAnswer("Agree")}
                >
                  👍 Agree
                </Button>
                <Button
                  variant="outline"
                  className="w-full sm:w-auto text-gray-700 border-gray-300 hover:bg-gray-100"
                  onClick={() => handleAnswer("Neutral")}
                >
                  😐 Neutral
                </Button>
                <Button
                  variant="outline"
                  className="w-full sm:w-auto text-red-700 border-red-400 hover:bg-red-100"
                  onClick={() => handleAnswer("Disagree")}
                >
                  👎 Disagree
                </Button>
              </div>
            </CardContent>
          </Card>
        </motion.div>

        <p className="text-center text-sm text-gray-500 mt-6">
          {progress}% complete
        </p>
      </div>
    </div>
  )
}
