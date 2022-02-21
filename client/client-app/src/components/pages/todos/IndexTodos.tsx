import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import PostCard from "../../organisms/posts/PostCard";
import useLoadingState from "../../../hooks/useLoadingState";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RooState } from "../../../reducks/store/store";
import { indexTodos } from "../../../reducks/todos/operations";
import { Todos } from "../../../reducks/todos/types";

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