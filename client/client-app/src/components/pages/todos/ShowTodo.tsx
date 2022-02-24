import { push } from "connected-react-router";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { Button, CardMedia, Checkbox, Container, FormGroup, Grid, Paper, Typography } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import sample1 from "../../../assets/images/sample1.jpeg"
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RooState } from "../../../reducks/store/store";
import { deleteTodo, showTodo } from "../../../reducks/todos/operations";
import useLoadingState from "../../../hooks/useLoadingState";
import { FormControlLabel } from "@material-ui/core";

type Params = {
  id: string | undefined
}

export const ShowTodo: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RooState) => state.todo)
  const loadingState = useLoadingState()
  const [finish, setFinish] = useState(false)
  
  useEffect(() => {
    dispatch(showTodo(id))
    setFinish(todo.isFinished)
  }, [])

  const imagePath = `http://localhost:8000/api/img/${todo.imagePath}`
  
  const onClickToEdit = useCallback(() => {
    dispatch(push(`/todo/edit/${id}`))
  }, [id])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodo(id))
  }, [id])

  // const onChangeIsFinished = useCallback(() => {
  //   dispatch(updateIsFinished(id))
  // }, [id])

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
                justifyContent="center"
                spacing={4}
                marginTop="20px"
                marginBottom="20px"
              >
                <Grid 
                  item
                  xs={12}
                  sm={4}
                >
                  <CardMedia
                    component="img"
                    src={todo.imagePath? imagePath : sample1}
                    sx={{
                      boxShadow: 8
                    }}
                  />
                </Grid>
                <Grid
                  item
                  xs={12}
                  sm={8}
                >
                  <Typography
                    variant="h3"
                    sx={{
                      fontSize: {
                        xs: '20px',
                        md: '35px'
                      }
                    }}
                  >
                    Memo
                  </Typography>
                  <Typography>
                    {todo.content}
                  </Typography>
                </Grid>
              </Grid>
              <Paper
                sx={{
                  paddingY: {
                    sx: "10px",
                    md: "20px"
                  },
                  marginY: {
                    sx: "10px",
                    sm: "10px",
                    md: "20px"
                  },
                  paddingLeft: "10px"
                }}
              >
                <FormControlLabel 
                  control={<Checkbox 
                              checked={finish}
                              // onChange={onChangeIsFinished}
                            />} 
                  label="finish"
                />
                <Button
                  onClick={onClickToEdit}
                  sx={{
                    color: 'black'
                  }}
                >
                  <EditIcon />
                  Edit
                </Button>
                <Button
                  onClick={onClickDelete}
                  sx={{
                    color: 'black'
                  }}
                >
                  <DeleteIcon />
                  Delete
                </Button>
              </Paper>
            </Paper>
          </Container>
        )
      }
    </>
  )
}

export default ShowTodo;