import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RooState } from "../../../reducks/store/store";
import { indexTodos } from "../../../reducks/todos/operations";
import { Todo, Todos } from "../../../reducks/todos/types";

import PostCard from "../../organisms/posts/PostCard";

export const IndexTodos = () => {
  const dispatch = useDispatch()
  const todosObj: Todos = useSelector((state: RooState) => state.todos)
  // Object -> Array
  let todos = Object.values(todosObj)

  useEffect(() => {
    dispatch(indexTodos())
  }, [])


  return(
    <>
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
          Object.keys(todos).length > 0 && (
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
    </>
  )
}

export default IndexTodos;