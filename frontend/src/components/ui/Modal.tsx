// Modal.tsx

import React, { ReactNode, useEffect } from 'react';
import ReactDOM from 'react-dom';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
}

function Modal({ isOpen, onClose, children }: ModalProps) {
  useEffect(() => {
    // Close modal on Escape key press
    const handleEsc = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose();
      }
    };

    if (isOpen) {
      document.addEventListener('keydown', handleEsc);
    }

    return () => {
      document.removeEventListener('keydown', handleEsc);
    };
  }, [isOpen, onClose]);

  // Prevent background scrolling when modal is open
  useEffect(() => {
    if (isOpen) {
      document.body.style.overflow = 'hidden';
    }

    return () => {
      document.body.style.overflow = '';
    };
  }, [isOpen]);

  if (!isOpen) return null;

  return ReactDOM.createPortal(
    <div
      className="fixed inset-0 z-50 flex items-center justify-center overflow-auto bg-black bg-opacity-50"
      onClick={onClose} // Close modal when clicking on the backdrop
    >
      <div
        className="bg-white rounded-lg shadow-lg w-full max-w-2xl mx-auto relative"
        onClick={(e) => e.stopPropagation()} // Prevent closing when clicking inside the modal
      >
        {/* Close button */}
        <button
          className="absolute top-2 right-2 text-gray-600 hover:text-gray-800 text-2xl font-bold"
          onClick={onClose}
          aria-label="Close"
        >
          &times;
        </button>
        {children}
      </div>
    </div>,
    document.getElementById('modal-root') as HTMLElement
  );
}

interface ModalHeaderProps {
  children: React.ReactNode;
}

function ModalHeader({ children }: ModalHeaderProps) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      <h2 className="text-xl font-semibold">{children}</h2>
    </div>
  );
}

interface ModalContentProps {
  children: React.ReactNode;
}

function ModalContent({ children }: ModalContentProps) {
  return <div className="px-6 py-4">{children}</div>;
}

interface ModalFooterProps {
  children: React.ReactNode;
}

function ModalFooter({ children }: ModalFooterProps) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 flex justify-end space-x-2">
      {children}
    </div>
  );
}

export { Modal, ModalHeader, ModalContent, ModalFooter };
