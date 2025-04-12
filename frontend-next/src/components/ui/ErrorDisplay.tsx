import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

type ErrorDisplayProps = {
  error: string
  onRetry?: () => void
  className?: string
}

export function ErrorDisplay({ error, onRetry, className }: ErrorDisplayProps) {
  return (
    <div className={cn('flex flex-col items-center justify-center space-y-4', className)}>
      <div className="text-center space-y-2">
        <div className="text-6xl mb-4">😕</div>
        <h3 className="text-xl font-semibold text-gray-900">Oops! Something went wrong</h3>
        <p className="text-gray-600 max-w-md">{error}</p>
      </div>
      {onRetry && (
        <Button 
          onClick={onRetry}
          className="mt-4 bg-blue-500 hover:bg-blue-600 text-white"
        >
          Try Again
        </Button>
      )}
    </div>
  )
} 