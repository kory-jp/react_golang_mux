import { push } from "connected-react-router";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useLocation } from "react-router-dom";
import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import useReturnTop from "../../../hooks/useReturnTop";
import { RootState } from "../../../reducks/store/store";
import { search } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";
import DefaultIndexSection from "./DefaultIndexSection";

export const SearchIndexSection: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()
  const query =  new URLSearchParams(useLocation().search)
  const tagId = Number(query.get("tagId"))
  const importance = Number(query.get("importance"))
  const urgency = Number(query.get("urgency"))
  const returnTop = useReturnTop()

  useEffect(()=> {
    dispatch(search(tagId, importance, urgency, queryPage, setSumPage))
  },[tagId, importance, urgency, queryPage])

  const todos: Todos = useSelector((state: RootState) => state.todos)
  
  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo/search?tagId=${tagId}&importance=${importance}&urgency=${urgency}&page=${page}`))
    returnTop()
  }, [tagId, importance, urgency])

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

export default SearchIndexSection;