import { useState } from "react"

export function useProgress(max: number = 100) {
  const [progress, setProgress] = useState(0)

  const increment = (step = 10) => {
    setProgress((p) => Math.min(max, p + step))
  }

  return { progress, increment }
}
