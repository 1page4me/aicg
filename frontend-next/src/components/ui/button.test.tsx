import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import { Button, buttonVariants } from './button'

/**
 * Test suite for the Button component
 * Testing strategy:
 * 1. Basic rendering and default props
 * 2. Variant styles
 * 3. Size variations
 * 4. Interaction handling
 * 5. Component composition (asChild)
 * 6. Style customization
 * 7. Accessibility features
 */
describe('Button Component', () => {
  // Clear all mocks after each test to prevent any side effects between tests
  afterEach(() => {
    jest.clearAllMocks()
  })

  // Test default rendering with snapshot for regression testing
  it('renders with default props', () => {
    const { container } = render(<Button>Click me</Button>)
    const button = screen.getByRole('button', { name: /click me/i })
    expect(button).toBeInTheDocument()
    expect(button).toHaveClass('inline-flex')
    expect(container).toMatchSnapshot() // Captures DOM structure for regression testing
  })

  // Test all button variants using a table-driven approach
  // This reduces code duplication and makes it easy to add new variants
  it.each([
    ['default', 'bg-blue-600'],
    ['destructive', 'bg-red-600'],
    ['outline', 'border-gray-300'],
    ['secondary', 'bg-gray-200'],
    ['ghost', 'hover:bg-gray-100'],
    ['link', 'text-blue-600']
  ])('renders %s variant with correct styles', (variant, expectedClass) => {
    render(<Button variant={variant as any}>Button</Button>)
    const button = screen.getByRole('button', { name: /button/i })
    expect(button).toHaveClass(expectedClass)
  })

  // Test all button sizes using a table-driven approach
  // Tests both single and compound class names
  it.each([
    ['default', 'h-9 px-4'],
    ['sm', 'h-8'],
    ['lg', 'h-10'],
    ['icon', 'h-9 w-9']
  ])('renders %s size with correct styles', (size, expectedClass) => {
    render(<Button size={size as any}>Button</Button>)
    const button = screen.getByRole('button', { name: /button/i })
    // Split compound classes and check each individually
    expectedClass.split(' ').forEach(className => {
      expect(button).toHaveClass(className)
    })
  })

  // Test click event handling and callback invocation
  // Verifies both single and multiple click scenarios
  it('handles click events', () => {
    const handleClick = jest.fn()
    render(<Button onClick={handleClick}>Click me</Button>)
    const button = screen.getByRole('button', { name: /click me/i })
    
    fireEvent.click(button)
    expect(handleClick).toHaveBeenCalledTimes(1)
    
    fireEvent.click(button)
    expect(handleClick).toHaveBeenCalledTimes(2)
  })

  // Test component composition using asChild prop
  // Verifies that the Button can wrap other elements while maintaining styles
  it('renders as a child component when asChild is true', () => {
    const { container } = render(
      <Button asChild>
        <a href="#test">Link Button</a>
      </Button>
    )
    const link = screen.getByRole('link', { name: /link button/i })
    expect(link).toBeInTheDocument()
    expect(link).toHaveClass('inline-flex')
    expect(link).toHaveAttribute('href', '#test')
    expect(container).toMatchSnapshot()
  })

  // Test custom class name application while preserving base styles
  // Ensures custom styles don't override essential button functionality
  it('applies custom className while preserving base styles', () => {
    const { container } = render(
      <Button className="custom-class test-class">Button</Button>
    )
    const button = screen.getByRole('button', { name: /button/i })
    expect(button).toHaveClass('custom-class', 'test-class', 'inline-flex')
    expect(container).toMatchSnapshot()
  })

  // Test disabled state behavior and styling
  // Verifies both visual feedback and interaction prevention
  it('renders in disabled state with correct attributes and styles', () => {
    const handleClick = jest.fn()
    render(<Button disabled onClick={handleClick}>Disabled Button</Button>)
    const button = screen.getByRole('button', { name: /disabled button/i })
    
    expect(button).toBeDisabled()
    expect(button).toHaveClass('disabled:opacity-50')
    
    // Verify that click events are prevented when disabled
    fireEvent.click(button)
    expect(handleClick).not.toHaveBeenCalled()
  })

  // Test TypeScript prop type safety and HTML attribute passing
  // Ensures proper attribute forwarding and accessibility
  it('maintains type safety with TypeScript props', () => {
    const { container } = render(
      <Button
        type="submit"
        form="test-form"
        aria-label="Submit"
        data-testid="submit-button"
      >
        Submit
      </Button>
    )
    const button = screen.getByTestId('submit-button')
    expect(button).toHaveAttribute('type', 'submit')
    expect(button).toHaveAttribute('form', 'test-form')
    expect(button).toHaveAttribute('aria-label', 'Submit')
    expect(container).toMatchSnapshot()
  })
}) 