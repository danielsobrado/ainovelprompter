
import React, { createContext, useContext, useState, ReactNode } from 'react';

interface WailsReadyContextType {
  wailsReady: boolean;
  setWailsReady?: React.Dispatch<React.SetStateAction<boolean>>; // Optional: if App exclusively sets it
}

const WailsReadyContext = createContext<WailsReadyContextType>({
  wailsReady: false,
});

export const useWailsReady = () => useContext(WailsReadyContext);

// Props for a potential Provider component if we abstract App's logic, but App.tsx will provide directly.
// interface WailsReadyProviderProps {
//   children: ReactNode;
// }

// export const WailsReadyProvider: React.FC<WailsReadyProviderProps> = ({ children }) => {
//   const [wailsReady, setWailsReady] = useState(false);
//   // Logic to listen to Wails event would go here if this was a self-contained provider
//   return (
//     <WailsReadyContext.Provider value={{ wailsReady, setWailsReady }}>
//       {children}
//     </WailsReadyContext.Provider>
//   );
// };

export default WailsReadyContext;
