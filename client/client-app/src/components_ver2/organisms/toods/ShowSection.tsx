import { CardMedia, Divider, Grid } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import useLoadingState from "../../../hooks/useLoadingState";
import { RootState } from "../../../reducks/store/store";
import { Tags } from "../../../reducks/tags/types";
import { deleteTodo, showTodo, updateIsFinished } from "../../../reducks/todos/operations";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import { PrimaryChip } from "../../atoms/chip/PrimaryChip";
import DefaultImage from "../../../assets/images/DefaultImage.jpg"
import taskComment from "../../../assets/images/taskComment.svg"
import EditIconsArea from "../../molecules/iconsArea/EditIconsArea";
import TagSection from "../../molecules/tag/TagSection";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";

type Params = {
  id: string | undefined
}

export const ShowSection: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RootState) => state.todos[0])
  const loadingState = useLoadingState()
  const [finish, setFinish] = useState(false)
  const tags: Tags | null = todo.tags ? todo.tags : null

  useEffect(() => {
    dispatch(showTodo(id))
  }, [id])

  useEffect(()=> {
    setFinish(todo.isFinished)
  },[todo])

  const imagePath = process.env.REACT_APP_API_URL + `img/${todo.imagePath}`

  const onClickToEdit = useCallback(() => {
    // dispatch(push(`/todo/edit/${id}`))
    console.log("modal open!")
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

  // ------
  const onClickCreateTaskCard = useCallback(()=> {
    console.log('crete!')
  }, [])
  // -----

  return (
    <>
      <Box
        className='showContainer'
        bgcolor='#2D2A2A'
        sx={{
          borderRadius: '10px',
        }}
      >
        <Box
          className='show__innner'
          sx={{
            padding: {
              xs: '16px',
            }
          }}
        >
          <Box
            className='show__title'
            sx={{
              marginBottom: {
                xs: '8px',
              }
            }}
          >
            <Box>
              <Box
                component='h1'
                color='#FFF'
                sx={{
                  fontSize: '16px',
                }}
              >
                {todo.title}
              </Box>
              <Divider
                sx={{
                  backgroundColor: '#FFF',
                }}
              />
            </Box>
          </Box>
          <Box
            className='show__infoArea'
            display='flex'
            justifyContent='space-between'
            sx={{
              marginBottom: {
                xs: '16px',
              }
            }}
          >
            <Box
              className='infoBadge'
              display='flex'
            >
              <Box
                sx={{
                  marginRight: {
                    xs: '8px',
                  }
                }}
              >
                <PrimaryChip 
                  label='重要'
                  colorFlag={todo.importance}
                />
              </Box>
              <Box>
                <PrimaryChip 
                  label='緊急'
                  colorFlag={todo.urgency}
                />
              </Box>
            </Box>
            <Box
              className='createdAt'
              color='#FFF'
            >
              {handleToDateFormat(todo.createdAt)}
            </Box>
          </Box>
          <Box
            className='show__image'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            <CardMedia
              component="img"
              image={todo.imagePath? imagePath : DefaultImage}
              sx={{
                height : {
                  xs: 'auto'
                },
                width: {
                  xs: '100%',
                },
                transition: '0.7s',
                filter: finish? 'grayscale(100%)' : '',
                '&:hover': {
                  cursor: 'pointer'
                }
              }}
            />
          </Box>
          <Box
            className='show__content'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            <Box
              className='content__title'
              sx={{
                marginBottom: {
                  xs: '16px',
                }
              }}
            >
              <Box
                color='#FFF'
                sx={{
                  fontSize: {
                    xs: '16px',
                  }
                }}
              >
                補足
              </Box>
              <Divider
                sx={{
                  backgroundColor: '#FFF',
                }}
              />              
            </Box>
            <Box
              className='content__content'
              color='#FFF'
              sx={{
                fontSize: {
                  xs: '16px',
                },
                minHeight: {
                  xs: '160px',
                }
              }}
            >
              {todo.content}
            </Box>
          </Box>
          <Box
            className='show__iconMenu'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            <EditIconsArea 
              finish={finish}
              onChangeIsFinished={onChangeIsFinished}
              onClickDelete={onClickDelete}
              onClickToEditTodo={onClickToEdit}
            />
          </Box>
          <Box
            className='show__tags'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            <TagSection 
              tags={todo.tags}
              onClickToSearchTagTodo={onClickToSearchTagTodo}
            />
          </Box>
          <Box
            className='show__task'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            <Box
              sx={{
                marginBottom: {
                  xs: '16px',
                }
              }}
            >
              <CardMedia
                component="img"
                image={taskComment}
                sx={{
                  height : {
                    xs: 'auto'
                  },
                  width: {
                    xs: '320px',
                  },
                }}
              />
            </Box>
            <Box>
              <PrimaryButton
                onClick={onClickCreateTaskCard}
              >
                タスクカードを作成
              </PrimaryButton>
            </Box>
          </Box>
        </Box>
      </Box>
    </>
  )
}

export default ShowSection;