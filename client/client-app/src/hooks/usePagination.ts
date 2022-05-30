import React, { useState } from "react";
import { useLocation } from "react-router-dom";

type PaginationType = {
  (): {
    sumPage: number,
    setSumPage: React.Dispatch<React.SetStateAction<number>>,
    queryPage: number
  }
}

export const usePagination: PaginationType = () => {
  const [sumPage, setSumPage] = useState(1)
  const {search} = useLocation()
  const query = new URLSearchParams(search)
  const queryPageStr = query.get("page") ? query.get("page") : "1"
  const queryPage = Number(queryPageStr)
  return { sumPage, setSumPage, queryPage}
}

export default usePagination;