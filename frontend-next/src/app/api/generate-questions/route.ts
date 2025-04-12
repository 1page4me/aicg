import { NextResponse } from "next/server"
import { mockQuestions } from "./mockQuestions"

export async function GET() {
  return NextResponse.json({ questions: mockQuestions })
}
