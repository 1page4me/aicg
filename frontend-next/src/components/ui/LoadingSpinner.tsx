import { cn } from '@/lib/utils'

type LoadingSpinnerProps = {
  className?: string
  text?: string
}

export function LoadingSpinner({ className, text = 'Loading...' }: LoadingSpinnerProps) {
  return (
    <div className={cn('flex flex-col items-center justify-center', className)}>
      <div className="relative">
        <div className="w-12 h-12 rounded-full border-4 border-gray-200"></div>
        <div className="w-12 h-12 rounded-full border-4 border-t-blue-500 border-r-blue-500 animate-spin absolute top-0 left-0"></div>
      </div>
      <p className="mt-4 text-gray-600 animate-pulse">{text}</p>
    </div>
  )
} 