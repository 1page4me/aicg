export interface UserProfile {
    name: string
    age: number
    country: string
    intelligenceProfile: IntelligenceTypePercentages
    aptitudes: AptitudeTypePercentages
  }
  
  export interface IntelligenceTypePercentages {
    logical: number
    linguistic: number
    interpersonal: number
    intrapersonal: number
    spatial: number
    musical: number
    bodilyKinesthetic: number
    naturalist: number
  }
  
  export interface AptitudeTypePercentages {
    design: number
    research: number
    quantitative: number
    mechanical: number
    verbal: number
    memory: number
  }
  