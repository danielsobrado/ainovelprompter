import React from 'react'
import { Button } from "./ui/button"
import { Settings } from 'lucide-react'
import StorageIndicator from './StorageIndicator'

interface HeaderProps {
  onSettingsClick: () => void;
}

export default function Header({ onSettingsClick }: HeaderProps) {
  return (
    <header className="flex justify-between items-center mb-4 bg-white py-2 px-4 border-b">
      <div className="flex items-center space-x-2">
        <h1 className="text-xl font-semibold">AI Novel Prompter</h1>
        <span className="text-xs text-gray-500">v0.1.0</span>
      </div>
      <div className="flex items-center gap-4">
        <StorageIndicator />
        <Button variant="ghost" size="icon" onClick={onSettingsClick}>
          <Settings className="h-5 w-5" />
        </Button>
      </div>
    </header>
  )
}