'use client'

import { Button } from '@/components/ui/button'
import { useRouter } from 'next/navigation'

export default function Home() {
  const router = useRouter()

  return (
    <main className="min-h-screen bg-gradient-to-b from-indigo-900 via-purple-900 to-pink-900 text-white">
      {/* Hero Section */}
      <div className="container mx-auto px-4 py-16">
        <div className="flex flex-col md:flex-row items-center justify-between gap-8">
          <div className="flex-1 space-y-6">
            <h1 className="text-5xl font-bold leading-tight">
              Master Your Knowledge with <span className="text-pink-400">AI-Powered</span> Quizzes
            </h1>
            <p className="text-xl text-gray-300">
              Challenge yourself with our interactive quizzes. Learn, compete, and track your progress in a fun way!
            </p>
            <div className="flex gap-4">
              <Button 
                size="lg" 
                className="bg-pink-500 hover:bg-pink-600 text-lg animate-bounce"
                onClick={() => router.push('/quiz')}
              >
                Start Quiz
              </Button>
              <Button 
                size="lg" 
                variant="outline" 
                className="border-white text-white hover:bg-white/10 text-lg"
                onClick={() => router.push('/leaderboard')}
              >
                View Leaderboard
              </Button>
            </div>
          </div>
          <div className="flex-1 relative">
            <div className="absolute inset-0 flex items-center justify-center">
              <div className="w-64 h-64 relative">
                {/* Animated circles */}
                <div data-testid="animated-circle" className="absolute inset-0 rounded-full border-4 border-pink-500 animate-ping"></div>
                <div data-testid="animated-circle" className="absolute inset-0 rounded-full border-4 border-purple-500 animate-ping" style={{ animationDelay: '0.2s' }}></div>
                <div data-testid="animated-circle" className="absolute inset-0 rounded-full border-4 border-indigo-500 animate-ping" style={{ animationDelay: '0.4s' }}></div>
                
                {/* Quiz icon */}
                <div className="absolute inset-0 flex items-center justify-center">
                  <svg className="w-32 h-32 text-white animate-bounce" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Features Section */}
      <div className="bg-black/10 backdrop-blur-sm py-16">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-bold text-center mb-12">Why Choose Our Platform?</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div data-testid="feature-card" className="bg-white/5 p-6 rounded-lg backdrop-blur-sm hover:scale-105 transition-transform duration-300">
              <div className="text-4xl mb-4 animate-bounce">🎯</div>
              <h3 className="text-xl font-semibold mb-2">Personalized Learning</h3>
              <p className="text-gray-300">Get quizzes tailored to your knowledge level and interests.</p>
            </div>
            <div data-testid="feature-card" className="bg-white/5 p-6 rounded-lg backdrop-blur-sm hover:scale-105 transition-transform duration-300">
              <div className="text-4xl mb-4 animate-bounce" style={{ animationDelay: '0.2s' }}>🏆</div>
              <h3 className="text-xl font-semibold mb-2">Compete & Win</h3>
              <p className="text-gray-300">Challenge friends and climb the leaderboard.</p>
            </div>
            <div data-testid="feature-card" className="bg-white/5 p-6 rounded-lg backdrop-blur-sm hover:scale-105 transition-transform duration-300">
              <div className="text-4xl mb-4 animate-bounce" style={{ animationDelay: '0.4s' }}>📊</div>
              <h3 className="text-xl font-semibold mb-2">Track Progress</h3>
              <p className="text-gray-300">Monitor your improvement with detailed analytics.</p>
            </div>
          </div>
        </div>
      </div>

      {/* CTA Section */}
      <div className="container mx-auto px-4 py-16 text-center">
        <h2 className="text-3xl font-bold mb-6">Ready to Test Your Knowledge?</h2>
        <p className="text-xl text-gray-300 mb-8">Join thousands of students who are already learning with us!</p>
        <Button 
          size="lg" 
          className="bg-pink-500 hover:bg-pink-600 text-lg px-8 animate-pulse"
          onClick={() => router.push('/register')}
        >
          Get Started Now
        </Button>
      </div>
    </main>
  )
}
