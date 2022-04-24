import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import PostCard from "../../organisms/posts/PostCard";
import useLoadingState from "../../../hooks/useLoadingState";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RootState } from "../../../reducks/store/store";
import { search } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";
import DefaultPagination from "../../molecules/pagination/DefaultPagination";
import usePagination from "../../../hooks/usePagination";
import { push } from "connected-react-router";
import { useLocation, useParams } from "react-router-dom";

type Params = {
  id: string | undefined
}

export const SearchByTagTodos: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()

  const query =  new URLSearchParams(useLocation().search)
  const tagId = Number(query.get("tagId"))
  const importance = Number(query.get("importance"))
  const urgency = Number(query.get("urgency"))

  useEffect(()=> {
    dispatch(search(tagId, importance, urgency, queryPage, setSumPage))
  },[tagId, importance, urgency, queryPage])

  const todos: Todos = useSelector((state: RootState) => state.todos)
  
  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo/search?tagId=${tagId}&importance=${importance}&urgency=${urgency}&page=${page}`))
  }, [tagId, importance, urgency])

  return(
    <>
     {
       loadingState? (
        <LoadingLayout />
       ) : (
        <Box
          id="index_content"
          sx={{
            marginY: {
              xs: '40px',
              md: '80px'
            },
            marginX: {
              xs: '10px',
              md: '50px'
            }
          }}
        >
          <Box
            sx={{
              marginBottom: {
                xs: '40px',
                md: '60px'
              }
            }}
          >
          </Box>
          {
            todos != null && todos.length > 0 && (
              <Grid container rowSpacing={1} columnSpacing={{xs: 1, sm: 2, md: 3}} paddingX={2}>
                {
                  todos.map(todo => (
                    <Grid 
                      item 
                      key={todo.id}
                      xs={12}
                      sm={6}
                      md={3}
                    > 
                      <PostCard 
                        todo={todo}
                        setSumPage={setSumPage}
                        queryPage={queryPage}
                      />
                    </Grid>
                  ))
                }
              </Grid>
            )
          } 
          <DefaultPagination 
            count={sumPage}
            onChange={onChangeCurrentPage}
            page={queryPage}
          />
        </Box>
       )
     }
    </>
  )
}

export default SearchByTagTodos;