import React, { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { ReadSettingsFile, WriteSettingsFile } from '../../wailsjs/go/main/App';

interface SettingsProps {
  isOpen: boolean;
  onClose: () => void;
}

interface Settings {
  maxFileSize: number;
  defaultLanguage: string;
  enableAutoSave: boolean;
  theme: 'light' | 'dark';
}

export function SettingsModal({ isOpen, onClose }: SettingsProps) {
  const [settings, setSettings] = useState<Settings>({
    maxFileSize: 500,
    defaultLanguage: 'en',
    enableAutoSave: true,
    theme: 'light',
  });

  useEffect(() => {
    if (isOpen) {
      loadSettings();
    }
  }, [isOpen]);

  const loadSettings = async () => {
    try {
      const loadedSettings = await ReadSettingsFile();
      if (loadedSettings) {
        setSettings(JSON.parse(loadedSettings));
      }
    } catch (error) {
      console.error("Error loading settings:", error);
    }
  };

  const saveSettings = async () => {
    try {
      await WriteSettingsFile(JSON.stringify(settings, null, 2));
      onClose();
    } catch (error) {
      console.error("Error saving settings:", error);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Settings</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="maxFileSize" className="text-right">
              Max File Size (KB)
            </Label>
            <Input
              id="maxFileSize"
              type="number"
              value={settings.maxFileSize}
              onChange={(e) => setSettings({ ...settings, maxFileSize: Number(e.target.value) })}
              className="col-span-3"
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="defaultLanguage" className="text-right">
              Default Language
            </Label>
            <Input
              id="defaultLanguage"
              value={settings.defaultLanguage}
              onChange={(e) => setSettings({ ...settings, defaultLanguage: e.target.value })}
              className="col-span-3"
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="enableAutoSave" className="text-right">
              Enable Auto Save
            </Label>
            <Checkbox
              id="enableAutoSave"
              checked={settings.enableAutoSave}
              onCheckedChange={(checked) => setSettings({ ...settings, enableAutoSave: checked as boolean })}
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="theme" className="text-right">
              Theme
            </Label>
            <Select
              value={settings.theme}
              onValueChange={(value) => setSettings({ ...settings, theme: value as 'light' | 'dark' })}
            >
              <SelectTrigger className="col-span-3">
                <SelectValue placeholder="Select a theme" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="light">Light</SelectItem>
                <SelectItem value="dark">Dark</SelectItem>
              </SelectContent>
            </Select>
          </div>
        </div>
        <DialogFooter>
          <Button onClick={saveSettings}>Save Changes</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}