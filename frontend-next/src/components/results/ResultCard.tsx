import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'

type ResultCardProps = {
  result: {
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
}

export function ResultCard({ result }: ResultCardProps) {
  return (
    <Card className="hover:shadow-lg transition-shadow transform hover:scale-[1.02] duration-200">
      <CardHeader>
        <CardTitle className="flex justify-between items-center">
          <span className="text-lg font-bold">Quiz #{result.quiz_id}</span>
          <span className={`text-sm px-3 py-1 rounded-full ${
            result.is_passed 
              ? 'bg-green-100 text-green-800 animate-pulse' 
              : 'bg-red-100 text-red-800'
          }`}>
            {result.is_passed ? '🎉 Passed' : '😢 Failed'}
          </span>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          <div>
            <div className="flex justify-between text-sm text-gray-600 mb-1">
              <span>Score</span>
              <span className="font-bold">{result.score.toFixed(1)}%</span>
            </div>
            <div className="relative">
              <Progress 
                value={result.score} 
                className="h-2 bg-gray-100"
              />
              <div 
                className={`absolute top-0 left-0 h-2 rounded-full ${
                  result.is_passed ? 'bg-green-500' : 'bg-red-500'
                }`}
                style={{ width: `${result.score}%` }}
              />
            </div>
          </div>
          
          <div className="grid grid-cols-2 gap-4 text-sm">
            <div className="bg-gray-50 p-3 rounded-lg">
              <p className="text-gray-600">Correct Answers</p>
              <p className="font-medium text-lg">{result.correct_answers}/{result.total_questions}</p>
            </div>
            <div className="bg-gray-50 p-3 rounded-lg">
              <p className="text-gray-600">Time Taken</p>
              <p className="font-medium text-lg">
                {Math.floor(result.time_taken / 60)}m {result.time_taken % 60}s
              </p>
            </div>
          </div>
          
          <div className="bg-gray-50 p-3 rounded-lg">
            <p className="text-gray-600">Passing Score</p>
            <p className="font-medium text-lg">{result.passing_score}%</p>
          </div>
          
          <div className="text-sm text-gray-500 border-t pt-2">
            <p>Completed on {new Date(result.created_at).toLocaleDateString()}</p>
          </div>
        </div>
      </CardContent>
    </Card>
  )
} 