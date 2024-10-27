import React from 'react';
import { Button } from "@/components/ui/button";
import { Plus, Folder, Trash } from "lucide-react";

interface FileOperationsProps {
  onAddFile: () => void;
  onAddFolder: () => void;
  onClearAll: () => void;
}

export const FileOperations: React.FC<FileOperationsProps> = ({
  onAddFile,
  onAddFolder,
  onClearAll
}) => {
  return (
    <div className="mt-2 space-x-2">
      <Button variant="outline" onClick={onAddFile}>
        <Plus size={16} className="mr-2" /> Add File
      </Button>
      <Button variant="outline" onClick={onAddFolder}>
        <Folder size={16} className="mr-2" /> Add Folder
      </Button>
      <Button variant="outline" onClick={onClearAll}>
        <Trash size={16} className="mr-2" /> Clear All
      </Button>
    </div>
  );
};