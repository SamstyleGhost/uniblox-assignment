import { createContext, useContext, useState } from "react";

interface UserContextType {
  customerID: string | null;
  setCustomerID: React.Dispatch<React.SetStateAction<string | null>>;
}

interface UserContextProviderProps {
  children: React.ReactNode;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserContextProvider = ({ children }: UserContextProviderProps) => {
  
  const [customerID, setCustomerID] = useState<string | null>(null);

  console.log(customerID);
  

  return (
    <UserContext.Provider
      value={{
        customerID,
        setCustomerID
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

export const useUserContext = (): UserContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserContextProvider");
  }
  return context;
};