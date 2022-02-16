import { Box } from "@mui/system";
import { FC, useEffect } from "react";
import { Grid } from "@mui/material";
import { useDispatch, useSelector } from "react-redux";

import { indexTodos } from "../../../reducks/todos/operations";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import PostCard from "../../organisms/posts/PostCard";
import { RooState } from "../../../reducks/store/store";
import { Todos } from "../../../reducks/todos/types";
import useLoadingState from "../../../hooks/useLoadingState";

export const IndexTodos: FC = () => {
  const dispatch = useDispatch()
  const loadingState = useLoadingState()
  const todos: Todos = useSelector((state: RooState) => state.todos)

  useEffect(() => {
    dispatch(indexTodos())
  }, [])

  return(
    <>
     {
       loadingState? (
        <LoadingLayout />
       ) : (
        <Box
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
          {
            todos.length > 0 && (
              <Grid container rowSpacing={1} columnSpacing={{xs: 1, sm: 2, md: 3}} paddingX={2}>
                {
                  todos.map(todo => (
                    <Grid item key={todo.id}> 
                      <PostCard todo={todo}/>
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

export default IndexTodos;