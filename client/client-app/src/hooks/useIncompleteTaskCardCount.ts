import { useContext } from "react"
import { IncompleteTaskCardCountContext, IncompleteTaskCardCountType } from "../providers/IncompleteTaskCardCount"

export const useIncompleteTaskCardCount = (): IncompleteTaskCardCountType  => useContext(IncompleteTaskCardCountContext)