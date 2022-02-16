import { CardMedia, Container, Grid, Paper, Typography } from "@mui/material";
import { FC, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RooState } from "../../../reducks/store/store";
import sample1 from "../../../assets/images/sample1.jpeg"
import { showTodo } from "../../../reducks/todos/operations";
import useLoadingState from "../../../hooks/useLoadingState";

type Params = {
  id: string | undefined
}

export const ShowTodo: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todos = useSelector((state: RooState) => state.todos)
  const loadingState = useLoadingState()
  
  useEffect(() => {
    dispatch(showTodo(id))
  }, [])

  let todo = Object.fromEntries(
    Object.entries(todos).map(([key, value]) => [key, value])
  )

  return(
    <>
      {
        loadingState? (
          <LoadingLayout />
        ) : (
          <Container maxWidth='lg'>
            <Paper
              sx={{
                padding: {
                  xs: '5px',
                  md: '20px'
                },
                marginTop: '30px'
              }}
            >
              <Paper
                sx={{
                  padding: {
                    xs: '5px',
                    md: '15px'
                  }
                }}
                elevation={5}
              >
                <Typography
                  variant="h2"
                  fontWeight='bold'
                  sx={{
                    fontSize: {
                      xs: '25px',
                      md: '40px'
                    }
                  }}
                >
                  {todo.title}
                </Typography>
              </Paper>
              <Grid 
                container
                spacing={3}
                sx={{
                  marginTop: {
                    xs: '20px',
                    md: '50px'
                  }
                }}
              > 
                  <Grid item>
                    <CardMedia
                      component="img"
                      src={sample1}
                      sx={{
                        height: {
                          xs: '200px',
                          md: '350px'
                        },
                        boxShadow: 8
                      }}
                    />
                  </Grid>
                  <Grid item>
                    <Typography
                      variant="h3"
                      sx={{
                        fontSize: {
                          xs: '20px',
                          md: '35px'
                        }
                      }}
                    >
                      show momo
                    </Typography>
                    <Typography>
                      {todo.content}
                    </Typography>
                  </Grid>
                </Grid>
            </Paper>
          </Container>
        )
      }
    </>
  )
}

export default ShowTodo;