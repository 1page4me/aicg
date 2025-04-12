"use client"

import React, { useState } from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"

export default function OnboardingCard() {
  const [progress, setProgress] = useState(0)
  const router = useRouter()

  const handleStart = () => {
    if (progress < 100) {
      setProgress((prev) => prev + 20)
    } else {
      router.push("/quiz")
    }
  }

  return (
    <Card className="w-full max-w-md p-6 shadow-2xl rounded-2xl border border-gray-200">
      <CardContent className="space-y-6">
        <h1 className="text-3xl font-bold text-center text-gray-800">
          🔍 Start Your Ikigai Discovery
        </h1>
        <p className="text-center text-gray-600">
          Explore your strengths, values, and passions to find your perfect career path.
        </p>

        <Progress value={progress} className="h-4" />

        <div className="flex justify-center">
          <Button onClick={handleStart} className="w-full text-lg py-6">
            {progress < 100 ? "Begin Journey" : "Start Quiz"}
          </Button>
        </div>

        {progress > 0 && (
          <p className="text-center text-sm text-gray-500">
            You’re {progress}% ready to begin your career path journey 🔓
          </p>
        )}
      </CardContent>
    </Card>
  )
}
