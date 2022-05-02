import { push } from "connected-react-router";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { Box, Button, CardMedia, Checkbox, Chip, Container, Grid, Paper, Typography } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import sample1 from "../../../assets/images/sample1.jpeg"
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import { RootState } from "../../../reducks/store/store";
import { deleteTodo, showTodo, updateIsFinished } from "../../../reducks/todos/operations";
import useLoadingState from "../../../hooks/useLoadingState";
import { FormControlLabel } from "@material-ui/core";
import { Tags } from "../../../reducks/tags/types";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import InputTaskCardForm from "../../organisms/taskCards/InputTaskCardForm";
import { createTaskCard, indexTaskCards } from "../../../reducks/taskCards/operations";
import usePagination from "../../../hooks/usePagination";
import { TaskCards } from "../../../reducks/taskCards/types";
import IndexTaskCard from "../../organisms/taskCards/IndexTaskCard";

type Params = {
  id: string | undefined
}

export const ShowTodo: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RootState) => state.todos[0])
  const loadingState = useLoadingState()
  const [finish, setFinish] = useState(false)
  const tags: Tags | null = todo.tags ? todo.tags : null
  const [title, setTitle] = useState("")
  const [purpose, setPurpose] = useState("")
  const [content, setContent] = useState("")
  const [memo, setMemo] = useState("")
  const {sumPage, setSumPage, queryPage} = usePagination()

  useEffect(() => {
    dispatch(showTodo(id))
    dispatch(indexTaskCards(id, setSumPage, queryPage))
  }, [id])

  const taskCards: TaskCards = useSelector((state: RootState) => state.taskCards)

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

  // --- todo isFinished ----
  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(id, true))
    }
  }, [id, finish])

  const onClickToSearchTagTodo = useCallback((tagId: number) => {
    dispatch(push(`/todo/tag/${tagId}`))
  },[])

  // --- モーダル操作 ---
  const [open, setOpen] = useState(false);
  const onClickModalOpen = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickModalClose = useCallback(() => {
    setOpen(false)
  }, [])

  const onChangeTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])

  const onChangePurpose = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPurpose(event.target.value)
  },[setPurpose])

  const onChangeContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])

  const onChangeMemo = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setMemo(event.target.value)
  },[setMemo])

  const onClickNewTaskCard = useCallback(() => {
    dispatch(createTaskCard(id, title, purpose, content, memo, setSumPage, queryPage))
    setOpen(false)
  }, [id, title, purpose, content, memo])


  return(
    <>
      {
        loadingState? (
          <LoadingLayout />
        ) : (
          <>
          <Container maxWidth='lg'>
            <Paper
              sx={{
                transition: '0.7s',
                bgcolor: finish? 'text.disabled' : 'white',
                marginY: '30px',
                padding: {
                  xs: '5px',
                  md: '20px'
                },
              }}
            >
              {/* ------ タイトル ------ */}
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
                    marginBottom="20px"
                  >
                    {todo.title}
                  </Box>
                </Typography>
                <Box>
                  <Box
                  display="flex"
                  >
                    <Box
                      marginRight="30px"
                    >
                      重要度
                    </Box>
                    <Box>
                      {todo.importance === 1 ? "重要" : "重要ではない"}
                    </Box>
                  </Box>
                  <Box
                    display="flex"
                  >
                    <Box
                      marginRight="30px"
                    >
                      緊急度
                    </Box>
                    <Box>
                      {todo.urgency === 1 ? "緊急" : "緊急ではない"}
                    </Box>
                  </Box>
                </Box>
              </Paper>
              {/* ----- 画像、コンテンツ ----- */}
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
              {/* ----- 	編集セクション ----- */}
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
              {/* ----- タグセクション ----- */}
              {
                tags != null && (
                  <Paper
                  id="tag_section"
                  sx={{
                    paddingY: {
                      // sx: "10px",
                      md: "20px"
                    },
                    marginY: {
                      sx: "10px",
                      sm: "10px",
                      md: "20px"
                    },
                  }}
                >
                  <Grid container>
                    {
                      tags.map(tag => (
                        <Grid
                          key={tag.id}
                          marginLeft="10px"
                          marginY="20px"
                        >
                          <Chip 
                            label={tag.label}
                            onClick={() => onClickToSearchTagTodo(tag.id)}
                          />
                        </Grid>
                      ))
                    }
                  </Grid>
                </Paper>
                )
              } 

              {/* --- タスクカード新規作成 --- */}
              <Paper
                id="taskCard"
                sx={{
                  paddingY: {
                    // sx: "10px",
                    md: "20px"
                  },
                  marginY: {
                    sx: "10px",
                    sm: "10px",
                    md: "20px"
                  },
                }}
              >
                <Box
                  margin="20px"
                >
                  <PrimaryButton
                    onClick={onClickModalOpen}
                  >
                    タスクカード新規作成
                  </PrimaryButton>
                </Box>
              </Paper>
            </Paper>

            {/* --- taskCards index --- */}
            {
              taskCards !== null && taskCards.length > 0 && (
                <Box
                sx={{
                  marginBottom: "80px",
                }}
                >
                   {
                     taskCards.map(taskCard =>(
                       <IndexTaskCard 
                       key={taskCard.id}
                       taskCard={taskCard}
                       />
                       ))
                      }
                </Box>
              )
            }

          </Container>
          {/* --- モーダル --- */}
          <InputTaskCardForm 
            open={open}
            onClose={onClickModalClose}
            title={title}
            purpose={purpose}
            content={content}
            memo={memo}
            onChangeTitle={onChangeTitle}
            onChangePurpose={onChangePurpose}
            onChangeContent={onChangeContent}
            onChangeMemo={onChangeMemo}
            onClickNewTaskCard={onClickNewTaskCard}
          />
          </>
        )
      }
    </>
  )
}

export default ShowTodo;