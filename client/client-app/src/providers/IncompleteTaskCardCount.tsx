import { createContext, Dispatch, ReactNode, SetStateAction, useState } from "react";

export type IncompleteTaskCardCountType = {
  incompleteTaskCardCount: number;
  setIncompleteTaskCardCount: Dispatch<SetStateAction<number>>;
}

export const IncompleteTaskCardCountContext =  createContext<IncompleteTaskCardCountType>({} as IncompleteTaskCardCountType)

export const IncompleteTaskCardCountProvider = (props: {children: ReactNode}) => {
  const { children } = props
  const [incompleteTaskCardCount, setIncompleteTaskCardCount] = useState<number>(0)
  return (
    <IncompleteTaskCardCountContext.Provider
      value={{incompleteTaskCardCount, setIncompleteTaskCardCount}}
    >
      {children}
    </IncompleteTaskCardCountContext.Provider>
  )
}