import { CardMedia, Divider, Skeleton } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import { RootState } from "../../../reducks/store/store";
import { Tags } from "../../../reducks/tags/types";
import { deleteTodo, showTodo, updateIsFinished } from "../../../reducks/todos/operations";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import { PrimaryChip } from "../../atoms/chip/PrimaryChip";
import DefaultImage from "../../../assets/images/DefaultImage.jpg"
import EditIconsArea from "../../molecules/iconsArea/EditIconsArea";
import TagSection from "../../molecules/tag/TagSection";
import EditTodoModal from "./EditTodoModal";
import useReturnTop from "../../../hooks/useReturnTop";
import TextFormat from "../../../utils/TextFormat";
import useSkeletonState from "../../../hooks/useSkeletonState";

type Params = {
  id: string | undefined
}

export const ShowSection: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const todo = useSelector((state: RootState) => state.todos[0])
  const [finish, setFinish] = useState(false)
  const tags: Tags | null = todo.tags ? todo.tags : null
  const [openModal ,setOpenModal] = useState(false)
  const returnTop = useReturnTop()
  const skeletonState = useSkeletonState()

  useEffect(() => {
    dispatch(showTodo(id))
  }, [dispatch, id])

  useEffect(()=> {
    setFinish(todo.isFinished)
  },[todo])

  let imagePath: string = ""
  if (process.env.NODE_ENV === "production") {
    imagePath = todo.imagePath ? todo.imagePath : ""
  } else {
    imagePath = todo.imagePath ? process.env.REACT_APP_API_URL + `img/${todo.imagePath}` : ""
  }

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodo(id))
  }, [dispatch, id])

  // --- todo isFinished ----
  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(id, true))
    }
  }, [dispatch, id, finish])

  const onClickToSearchTagTodo = useCallback((tagId: number) => {
    dispatch(push(`/todo/search?tagId=${tagId}&importance=0&urgency=0&page=1`))
    returnTop()
  },[dispatch, returnTop])

  const onClickOpenEditTodoModal = useCallback(() => {
    setOpenModal(true)
  }, [])

  const onClickCloseTodoModal = useCallback(() => {
    setOpenModal(false)
  }, [])


  return (
    <>
      <Box
        className='showContainer'
        sx={{
          borderRadius: '10px',
          bgcolor: finish? '#464141' : '#2D2A2A',
          minWidth: {
            xs: '320px',
          }
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
                  fontSize: {
                    xs: '16px',
                    md: '24px',
                  }
                }}
              >
                {
                  skeletonState ? (
                    <Skeleton
                      variant="rectangular"
                      animation="wave"
                      sx={{
                        width: '100%',
                        height: '40px',
                      }}
                    />
                  ) : (
                    <>
                      {todo.title}
                    </>
                  )
                }
              </Box>
              <Divider
                sx={{
                  backgroundColor: '#FFF',
                  marginBottom: '16px',
                }}
              />
            </Box>
          </Box>
          {
            skeletonState ? (
              <Skeleton
                variant="rectangular" 
                animation='wave' 
                sx={{
                  width: '100%',
                  height: '24px',
                  marginBottom: '16px',
                }}                
              />
            ) : (
              <>
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
              </>
            )
          }
          <Box
            className='show__image'
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            {
              skeletonState ? (
                <Skeleton 
                  variant="rectangular" 
                  animation='wave' 
                  sx={{
                    width: '100%',
                    height: '300px',
                    marginBottom: '16px',
                  }} 
                />
              ) : (
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
              )
            }
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
              {
                skeletonState ? (
                  <Skeleton 
                    variant="rectangular" 
                    animation='wave' 
                    sx={{
                      width: '100%',
                      height: '80px',
                      marginBottom: '16px',
                    }}                      
                  />
                ) : (
                  <TextFormat 
                    text={todo.content}
                  />                  
                )
              }
            </Box>
          </Box>
           {
             skeletonState ? (
               <>
                <Skeleton
                  variant="rectangular" 
                  animation='wave' 
                  sx={{
                    width: '100%',
                    height: '24px',
                    marginBottom: '16px',
                  }}  
                />
                <Skeleton
                  variant="rectangular" 
                  animation='wave' 
                  sx={{
                    width: '100%',
                    height: '24px',
                    marginBottom: '16px',
                  }}  
                />                                
               </>
             ) : (
               <>
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
                    onClickToEdit={onClickOpenEditTodoModal}
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
                    tags={tags}
                    onClickToSearchTagTodo={onClickToSearchTagTodo}
                  />
                </Box>
               </>
             )
           }
        </Box>
      </Box>
      <EditTodoModal 
        open={openModal}
        onClose={onClickCloseTodoModal}
      />
    </>
  )
}

export default ShowSection;