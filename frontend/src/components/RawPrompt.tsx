import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"

interface RawPromptProps {
  value: string;
  onChange: (value: string) => void;
}

export default function RawPrompt({ value, onChange }: RawPromptProps) {
  return (
    <div className="mt-4">
      <Label htmlFor="raw-prompt" className="text-sm font-medium">
        Task Instruction (Raw Prompt)
      </Label>
      <Textarea
        id="raw-prompt"
        placeholder="Enter your task instruction here. Some examples:
- Add a button to calculate the total price of the items.
- Fix the memory leak in the search function.
- Refactor the state logic into a separate file."
        value={value}
        onChange={(e) => onChange(e.target.value)}
        className="mt-1 h-32"
      />
    </div>
  )
}