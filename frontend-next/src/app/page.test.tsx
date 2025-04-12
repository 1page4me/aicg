import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import Home from './page'
import { useRouter } from 'next/navigation'

// Mock the useRouter hook
jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}))

describe('Home Page', () => {
  const mockRouter = {
    push: jest.fn(),
  }

  beforeEach(() => {
    (useRouter as jest.Mock).mockReturnValue(mockRouter)
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('renders the hero section with correct content', () => {
    render(<Home />)
    
    // Check main heading
    expect(screen.getByText(/Master Your Knowledge with/i)).toBeInTheDocument()
    expect(screen.getByText(/AI-Powered/i)).toBeInTheDocument()
    expect(screen.getByText(/Quizzes/i)).toBeInTheDocument()
    
    // Check subheading
    expect(screen.getByText(/Challenge yourself with our interactive quizzes/i)).toBeInTheDocument()
  })

  it('renders all feature cards', () => {
    render(<Home />)
    
    // Check feature headings
    expect(screen.getByText('Personalized Learning')).toBeInTheDocument()
    expect(screen.getByText('Compete & Win')).toBeInTheDocument()
    expect(screen.getByText('Track Progress')).toBeInTheDocument()
    
    // Check feature descriptions
    expect(screen.getByText(/Get quizzes tailored to your knowledge level/i)).toBeInTheDocument()
    expect(screen.getByText(/Challenge friends and climb the leaderboard/i)).toBeInTheDocument()
    expect(screen.getByText(/Monitor your improvement with detailed analytics/i)).toBeInTheDocument()
  })

  it('renders the CTA section', () => {
    render(<Home />)
    
    expect(screen.getByText('Ready to Test Your Knowledge?')).toBeInTheDocument()
    expect(screen.getByText(/Join thousands of students/i)).toBeInTheDocument()
  })

  it('navigates to correct routes when buttons are clicked', () => {
    render(<Home />)
    
    // Start Quiz button
    fireEvent.click(screen.getByText('Start Quiz'))
    expect(mockRouter.push).toHaveBeenCalledWith('/quiz')
    
    // View Leaderboard button
    fireEvent.click(screen.getByText('View Leaderboard'))
    expect(mockRouter.push).toHaveBeenCalledWith('/leaderboard')
    
    // Get Started Now button
    fireEvent.click(screen.getByText('Get Started Now'))
    expect(mockRouter.push).toHaveBeenCalledWith('/register')
  })

  it('renders the animated quiz icon', () => {
    render(<Home />)
    
    const svg = document.querySelector('svg')
    expect(svg).toBeInTheDocument()
    expect(svg).toHaveClass('w-32', 'h-32', 'text-white', 'animate-bounce')
  })

  it('renders animated circles in hero section', () => {
    render(<Home />)
    
    const circles = screen.getAllByTestId('animated-circle')
    expect(circles).toHaveLength(3)
    circles.forEach((circle) => {
      expect(circle).toHaveClass('animate-ping')
    })
  })

  it('applies hover animations to feature cards', () => {
    render(<Home />)
    
    const cards = screen.getAllByTestId('feature-card')
    expect(cards).toHaveLength(3)
    cards.forEach((card) => {
      expect(card).toHaveClass('hover:scale-105', 'transition-transform', 'duration-300')
    })
  })
}) 