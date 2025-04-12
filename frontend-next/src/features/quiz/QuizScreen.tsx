// src/features/quiz/QuizScreen.tsx

"use client"

import React, { useState } from "react"
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Progress } from "@/components/ui/progress"

const questions = [
  {
    id: 1,
    question: "I enjoy solving logical puzzles and brain teasers.",
    category: "logicalIntelligence",
  },
  {
    id: 2,
    question: "I can easily imagine and rotate objects in space.",
    category: "spatialIntelligence",
  },
  {
    id: 3,
    question: "I express myself well in writing or speech.",
    category: "verbalLinguistic",
  },
]

export default function QuizScreen({ onComplete }: { onComplete: () => void }) {
  const [current, setCurrent] = useState(0)
  const [answers, setAnswers] = useState<Record<number, number>>({})

  const handleAnswer = (value: number) => {
    setAnswers({ ...answers, [questions[current].id]: value })
    if (current + 1 < questions.length) {
      setCurrent(current + 1)
    } else {
      onComplete()
    }
  }

  const progressPercent = ((current + 1) / questions.length) * 100

  return (
    <Card className="w-full max-w-xl p-6 shadow-lg border rounded-xl">
      <CardContent className="space-y-6">
        <Progress value={progressPercent} />

        <h2 className="text-xl font-semibold text-gray-800">
          {questions[current].question}
        </h2>

        <div className="grid grid-cols-5 gap-2">
          {[1, 2, 3, 4, 5].map((value) => (
            <Button
              key={value}
              onClick={() => handleAnswer(value)}
              variant="outline"
              className="py-6"
            >
              {value}
            </Button>
          ))}
        </div>

        <p className="text-sm text-center text-gray-500">
          Rate from 1 (Strongly Disagree) to 5 (Strongly Agree)
        </p>
      </CardContent>
    </Card>
  )
}
