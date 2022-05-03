import { push } from "connected-react-router";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import useReturnTop from "../../../hooks/useReturnTop";
import { RootState } from "../../../reducks/store/store";
import { indexTodos } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";
import DefaultIndexSection from "./DefaultIndexSection";

export const IndexSection: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()
  const returnTop = useReturnTop()
  
  useEffect(() => {
    dispatch(indexTodos(setSumPage, queryPage))
  }, [setSumPage, queryPage])
  const todos: Todos = useSelector((state: RootState) => state.todos)

  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo?page=${page}`))
    returnTop()
  }, [])

  return(
    <>
      <DefaultIndexSection 
        todos={todos}
        sumPage={sumPage}
        setSumPage={setSumPage}
        queryPage={queryPage}
        loadingState={loadingState}
        onChangeCurrentPage={onChangeCurrentPage}
      />
    </>
  )
}

export default IndexSection;