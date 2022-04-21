import { push } from "connected-react-router";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { Box, Button, CardMedia, Checkbox, Container, Grid, Paper, Typography } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import sample1 from "../../../assets/images/sample1.jpeg"
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RootState } from "../../../reducks/store/store";
import { deleteTodo, showTodo, updateIsFinished } from "../../../reducks/todos/operations";
import useLoadingState from "../../../hooks/useLoadingState";
import { FormControlLabel } from "@material-ui/core";

type Params = {
  id: string | undefined
}

export const ShowTodo: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RootState) => state.todo)
  const loadingState = useLoadingState()
  const [finish, setFinish] = useState(false)

  useEffect(() => {
    dispatch(showTodo(id))
  }, [id])

  useEffect(()=> {
    setFinish(todo.isFinished)
  },[todo])

  const imagePath = process.env.REACT_APP_API_URL + `img/${todo.imagePath}`
  
  const onClickToEdit = useCallback(() => {
    dispatch(push(`/todo/edit/${id}`))
  }, [id])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodo(id))
  }, [id])

  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(id, true))
    }
  }, [id, finish])


  return(
    <>
      {
        loadingState? (
          <LoadingLayout />
        ) : (
          <Container maxWidth='lg'>
            <Paper
              sx={{
                transition: '0.7s',
                bgcolor: finish? 'text.disabled' : 'white',
                marginTop: '30px',
                padding: {
                  xs: '5px',
                  md: '20px'
                },
              }}
            >
              <Paper
                sx={{
                  transition: '0.7s',
                  bgcolor: finish? '#bdbdbd' : 'white',
                  padding: {
                    xs: '5px',
                    md: '15px'
                  }
                }}
                elevation={5}
              >
                <Typography
                  variant="h2"
                >
                  <Box
                    sx={{
                      fontSize: {
                        xs: '20px',
                        sm: '30px',
                        md: '40px'
                      },
                    }}
                    fontWeight="bold"
                  >
                    {todo.title}
                  </Box>
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
                  >
                    <Box
                      sx={{
                        fontSize: {
                          xs: '20px',
                          sm: '30px',
                          md: '40px'
                        },
                      }}                      
                    >
                      Memo
                    </Box>
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
                <Grid container>
                  <Grid 
                    item
                    sx={{
                      marginBottom: '5px'
                    }}
                  >
                    <FormControlLabel 
                      control={<Checkbox 
                                  checked={finish}
                                  value={finish}
                                  onChange={onChangeIsFinished}
                                />} 
                      label="finish"
                    />
                  </Grid>
                  <Grid 
                    item
                    sx={{
                      marginBottom: '5px'
                    }} 
                  >
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
                  </Grid>
                </Grid>
              </Paper>
            </Paper>
          </Container>
        )
      }
    </>
  )
}

export default ShowTodo;