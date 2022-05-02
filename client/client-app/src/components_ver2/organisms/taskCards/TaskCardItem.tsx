import { Box, CardActions, Divider, FormControlLabel, Grid } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import usePagination from "../../../hooks/usePagination";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import { TaskCard } from "../../../reducks/taskCards/types";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import MoreIconsArea from "../../molecules/iconsArea/MoreIconsArea";
import ShowTCModal from "./ShowTCModal";

type Props = {
  taskCard: TaskCard
}

export const TaskCardItem: FC<Props> = (props) => {
  const { taskCard } = props
  const dispatch = useDispatch()
  const [isFinished, setIsFinished] = useState(false)
  const [openShowTCModal, setOpenShowTCModal] = useState(false)
  const {setSumPage, queryPage} = usePagination()

    useEffect(() => {
    setIsFinished(taskCard.isFinished)
  },[taskCard.isFinished])

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
  }, [taskCard])

  const onClickOpenShowTCModal = useCallback(() => {
    setOpenShowTCModal(true)
  }, [])

  const onClickCloseShowTCModal = useCallback(() => {
    setOpenShowTCModal(false)
  }, [])

  return(
    <>
      <Box
        className="taskcardItem"
        sx={{
          backgroundColor: isFinished ? '#464141' : '#2D2A2A',
          borderRadius: "10px",
          padding: "20px",
        }}
      >
        <Box
          className="tc__item__inner"
        >
          <Box
            className="tc__item__inner_title"
            color='#FFF'
            sx={{
              fontSize: {
                xs: '16px',
                nd: '24px',
              },
              marginBottom: {
                xs: '8px',
              }
            }}
          >
            <Box
              sx={{
                marginBottom: {
                  xs: '16px'
                },
                ":hover": {
                  cursor: 'pointer',
                }
              }}
              onClick={onClickOpenShowTCModal}
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
            className="tc__item__inner__createdAt"
            textAlign='end'
            sx={{
              marginBottom: {
                xs: '24px',
              }
            }}
          >
            <Box
              sx={{
                color: '#FFF',
                fontSize: {
                  xs: '16px',
                }
              }}
            >
              {handleToDateFormat(taskCard.createdAt)}              
            </Box>
          </Box>
          <Box>
            <CardActions>
              <MoreIconsArea 
                finish={isFinished}
                onChangeIsFinished={onChangeIsFinished}
                onClickDelete={onClickDelete}
                onClickMoreInfo={onClickOpenShowTCModal}
              />
            </CardActions>
          </Box>
        </Box>
      </Box>
      <ShowTCModal 
        open={openShowTCModal}
        onClose={onClickCloseShowTCModal}
        taskCard={taskCard}
      />
    </>
  )
}

export default TaskCardItem;