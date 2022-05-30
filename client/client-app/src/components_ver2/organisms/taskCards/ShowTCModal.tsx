import { Divider, Modal } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useEffect, useState } from "react";
import CloseIcon from '@mui/icons-material/Close';
import { TaskCard } from "../../../reducks/taskCards/types";
import EditIconsArea from "../../molecules/iconsArea/EditIconsArea";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import { useDispatch } from "react-redux";
import usePagination from "../../../hooks/usePagination";
import EditTCModal from "./EditTCModal";
import TextFormat from "../../../utils/TextFormat";
import { useIncompleteTaskCardCount } from "../../../hooks/useIncompleteTaskCardCount";

type Props = {
  open: boolean,
  onClose: () => void,
  taskCard: TaskCard
}

export const ShowTCModal: FC<Props> = (props) => {
  const {open, onClose, taskCard} = props
  const dispatch = useDispatch()
  const [isFinished, setIsFinished] = useState(false)
  const {setSumPage, queryPage} = usePagination()
  const [openEditModal, setOpenEditModal] = useState(false)
  const {incompleteTaskCardCount, setIncompleteTaskCardCount} = useIncompleteTaskCardCount()

  useEffect(() => {
    setIsFinished(taskCard.isFinished)
  }, [taskCard.isFinished])

  const onChangeIsFinished = useCallback(() => {
    if (isFinished) {
      setIsFinished(false)
      dispatch(updateIsFinished(taskCard.id, false))
      taskCard.isFinished = false
      setIncompleteTaskCardCount(incompleteTaskCardCount + 1)
    } else {
      setIsFinished(true)
      dispatch(updateIsFinished(taskCard.id, true))
      taskCard.isFinished = true
      setIncompleteTaskCardCount(incompleteTaskCardCount - 1)
    }
  }, [isFinished, dispatch, incompleteTaskCardCount, setIncompleteTaskCardCount, taskCard])

  const onClickDelete = useCallback(() => {
    setTimeout(()=> {
      setIncompleteTaskCardCount(incompleteTaskCardCount - 1)
    }, 500)
    dispatch(deleteTaskCard(taskCard.id, taskCard.todoId, setSumPage, queryPage))
    onClose()
  }, [dispatch, incompleteTaskCardCount, onClose, queryPage, setIncompleteTaskCardCount, setSumPage, taskCard.id, taskCard.todoId])

  const onClickOpenEditModal = useCallback(() => {
    setOpenEditModal(true)
  }, [])

  const onClickCloseEditModal = useCallback(() => {
    setOpenEditModal(false)
  }, [])

  return(
    <>
      <Modal
        open={open}
        onClose={onClose}
        sx={{
          overflow: 'scroll',
        }}
      >
        <Box
          className='createTodoModal'
          sx={{
            backgroundColor: isFinished ? '#464141' : '#2D2A2A',
            marginX: 'auto',
            marginTop: '5%',
            width: {
              xs: '90%',
              sm: '70%',
              md: '60%',
              lg: '50%',
            },
            borderRadius: '10px',
          }}
        >
          <Box
            className='modal__inner'
            sx={{
              padding: {
                xs: '16px',
              }
            }}
          >
            <Box
              className='close'
              textAlign='end'
            >
              <Box
                className='close__button'
              >
                <CloseIcon
                  fontSize="large"
                  onClick={onClose}
                  sx={{
                    color: '#FFF',
                    cursor: 'pointer',
                  }}
                />
              </Box>
            </Box>
            <Box
              className='create_tc_title'
              component='h2'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}              
            >
              <Box
                sx={{
                  marginBottom: {
                    xs: '8px',
                  },
                  fontSize: {
                    xs: '16px',
                    md: '24px',
                  }
                }}
              >
                {taskCard.title}
              </Box>
              <Divider 
                sx={{
                  backgroundColor: '#FFF'
                }}
              />
            </Box>
            <Box
              className='tc_purpose'
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
                目的
              </Box>
              <Box
                sx={{
                  marginBottom: {
                    xs: '8px',
                  }
                }}
              >
                なぜこのタスクをする必要があるのか
              </Box>
              <Divider 
                sx={{
                  backgroundColor: '#FFF',
                  marginBottom: {
                    xs: '16px',
                  }
                }}
              />
              <Box
                sx={{
                  minHeight: {
                    xs: '80px',
                  }
                }}
              >
                <TextFormat 
                  text={taskCard.purpose}
                />
              </Box>            
            </Box>
            <Box
              className='tc__content'
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
                作業内容
              </Box>
              <Box
                sx={{
                  marginBottom: {
                    xs: '16px',
                  }
                }}
              >
                具体的にどのような作業をするのか
              </Box> 
              <Divider 
                sx={{
                  backgroundColor: '#FFF',
                  marginBottom: {
                    xs: '16px',
                  }
                }}
              />
              <Box
                sx={{
                  minHeight: {
                    xs: '80px',
                  }
                }}
              >
                <TextFormat 
                  text={taskCard.content}
                />
              </Box>                      
            </Box>
            <Box
              className='tc__memo'
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
                メモ
              </Box>
              <Divider 
                sx={{
                  backgroundColor: '#FFF',
                  marginBottom: {
                    xs: '16px',
                  }
                }}
              />              
              <Box
                sx={{
                  minHeight: {
                    xs: '80px',
                  }
                }}
              >
                <TextFormat 
                  text={taskCard.memo}
                />
              </Box>
            </Box>
            <Box
              className='icon__section'
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <EditIconsArea 
                finish={isFinished}
                onChangeIsFinished={onChangeIsFinished}
                onClickDelete={onClickDelete}
                onClickToEdit={onClickOpenEditModal}
              />
            </Box>
          </Box>
        </Box>
      </Modal>
      <EditTCModal 
        open={openEditModal}
        onClose={onClickCloseEditModal}
        taskCard={taskCard}
      />
    </>
  )
}

export default ShowTCModal;