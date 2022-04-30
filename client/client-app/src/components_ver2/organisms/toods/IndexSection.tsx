import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import { RootState } from "../../../reducks/store/store";
import { indexTodos } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import TodoCard from "./TodoCard";

export const IndexSection: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()
  
  useEffect(() => {
    dispatch(indexTodos(setSumPage, queryPage))
  }, [setSumPage, queryPage])
  const todos: Todos = useSelector((state: RootState) => state.todos)
  console.log(todos)

  return(
    <>
      {
        loadingState ? (
          <LoadingLayout />
        ) : (
          <Box>
            {
              todos != null && todos.length > 0 && (
                <Grid
                  container
                  // rowSpacing={1} 
                  // columnSpacing={{xs: 1, sm: 2, md: 3}} 
                  // paddingX={2}
                >
                  {
                    todos.map (todo => (
                      <Grid
                        item
                        key={todo.id}
                      >
                        <TodoCard 
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
          </Box>
        )
      }
    </>
  )
}

export default IndexSection;