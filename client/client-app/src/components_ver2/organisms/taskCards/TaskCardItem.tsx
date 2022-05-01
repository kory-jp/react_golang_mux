import { Box, CardActions, Divider, FormControlLabel, Grid } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import { TaskCard } from "../../../reducks/taskCards/types";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import MoreIconsArea from "../../molecules/iconsArea/MoreIconsArea";

type Props = {
  taskCard: TaskCard
}

export const TaskCardItem: FC<Props> = (props) => {
  const { taskCard } = props
  const dispatch = useDispatch()
  const [isFinished, setIsFinished] = useState(false)

    useEffect(() => {
    setIsFinished(taskCard.isFinished)
  },[])

  const onChangeIsFinished = useCallback(() => {
    if (isFinished) {
      setIsFinished(false)
      dispatch(updateIsFinished(taskCard.id, false))
    } else {
      setIsFinished(true)
      dispatch(updateIsFinished(taskCard.id, true))
    }
  }, [isFinished])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTaskCard(taskCard.id))
  }, [taskCard])

  // -----
  const onclickToShowTodo = useCallback(() => {
    console.log("open!")
    // setOpen(true)
  }, [])
  // ------

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
              fontSize: "24px",
              marginBottom: {
                xs: '8px',
              }
            }}
          >
            <Box
              sx={{
                marginBottom: {
                  xs: '16px'
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
                onClickToShowTodo={onclickToShowTodo}
              />
            </CardActions>
          </Box>
        </Box>
      </Box>
    </>
  )
}

export default TaskCardItem;