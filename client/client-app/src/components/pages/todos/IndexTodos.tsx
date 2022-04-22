import { FormControl, Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import PostCard from "../../organisms/posts/PostCard";
import useLoadingState from "../../../hooks/useLoadingState";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RootState } from "../../../reducks/store/store";
import { indexTodos } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";
import DefaultPagination from "../../molecules/pagination/DefaultPagination";
import usePagination from "../../../hooks/usePagination";
import { push } from "connected-react-router";
import TagSelction from "../../organisms/layout/TagSelction";
import { indexTags } from "../../../reducks/tags/operations";
import { Tags} from "../../../reducks/tags/types"

export const IndexTodos: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const todos: Todos = useSelector((state: RootState) => state.todos)
  const {sumPage, setSumPage, queryPage} = usePagination()

  useEffect(() => {
    dispatch(indexTodos(setSumPage, queryPage))
    dispatch(indexTags())
  }, [setSumPage, queryPage])

  const options = useSelector((state: RootState) => state.tags)

  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo?page=${page}`))
  }, [])
  const onChangeSelectTags = useCallback((event: React.SetStateAction<Tags>) => {
    console.log(event)
    // dispatch(searchTag(event))
  },[])


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
            <FormControl fullWidth>
              <TagSelction options={options} onChange={onChangeSelectTags}/>
            </FormControl>
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

export default IndexTodos;