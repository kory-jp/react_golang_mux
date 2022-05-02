import { Divider, Modal } from "@mui/material";
import { Box } from "@mui/system";
import { Dispatch, FC, useCallback, useEffect, useState } from "react";
import CloseIcon from '@mui/icons-material/Close';
import { TaskCard } from "../../../reducks/taskCards/types";
import EditIconsArea from "../../molecules/iconsArea/EditIconsArea";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import { useDispatch } from "react-redux";
import usePagination from "../../../hooks/usePagination";
import EditTCModal from "./EditTCModal";

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

  useEffect(() => {
    setIsFinished(taskCard.isFinished)
  }, [taskCard.isFinished])

  const onChangeIsFinished = useCallback(() => {
    if (isFinished) {
      setIsFinished(false)
      dispatch(updateIsFinished(taskCard.id, false))
      taskCard.isFinished = false
    } else {
      setIsFinished(true)
      dispatch(updateIsFinished(taskCard.id, true))
      taskCard.isFinished = true
    }
  }, [isFinished])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTaskCard(taskCard.id, taskCard.todoId, setSumPage, queryPage))
    onClose()
  }, [])

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
                onClick={onClose}
              >
                <CloseIcon
                  fontSize="large"
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
                {taskCard.purpose}
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
                {taskCard.content}
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
                {taskCard.memo}
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