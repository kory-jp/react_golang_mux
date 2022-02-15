import { CardMedia, Container, Grid, Paper, Typography } from "@mui/material";
import { FC, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useLocation, useParams } from "react-router-dom";
import sample1 from "../../../assets/images/sample1.jpeg"
import { RooState } from "../../../reducks/store/store";
import { showTodo } from "../../../reducks/todos/operations";

type Params = {
  id: string | undefined
}

export const ShowTodo: FC = () => {
  const location = useLocation()
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RooState) => state.todos)
  
  useEffect(() => {
    dispatch(showTodo(id))
  }, [])
  console.log(todo)


  return(
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
            show title
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
                show content
              </Typography>
            </Grid>
          </Grid>
      </Paper>
    </Container>
  )
}

export default ShowTodo;