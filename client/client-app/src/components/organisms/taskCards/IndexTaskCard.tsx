import { Box, Button, CardActions, Checkbox, FormControlLabel, Grid } from "@mui/material";
import { useCallback, useEffect, useState, VFC } from "react";
import DeleteIcon from '@mui/icons-material/Delete';
import StickyNote2Icon from '@mui/icons-material/StickyNote2';
import { TaskCard } from "../../../reducks/taskCards/types";
import ShowTaskCardModal from "./ShowTaskCardModal";
import { useDispatch } from "react-redux";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import handleToDateFormat from "../../../utils/handleToDateFormat";
import usePagination from "../../../hooks/usePagination";

type Props ={
  taskCard: TaskCard
}

export const IndexTaskCard: VFC<Props> = (props) => {
  const dispatch = useDispatch()
  const {taskCard} = props
  const [isFinished, setIsFinished] = useState(false)
  const [open, setOpen] = useState(false);
  const {setSumPage, queryPage} = usePagination()

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
    dispatch(deleteTaskCard(taskCard.id, taskCard.todoId, setSumPage, queryPage))
  }, [taskCard])

  const onclickToShowTodo = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickModalOpen = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickModalClose = useCallback(() => {
    setOpen(false)
  }, [])

  return(
    <>
      <Box
        sx={{
          backgroundColor: isFinished ? 'text.disabled' : 'white',
          borderRadius: "10px",
          padding: "20px",
          marginBottom: "40px",
        }}
      >
        <Box>
          <Box
            sx={{
              fontSize: "24px"
            }}
          >
            {taskCard.title}
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between"
            }}
          >
            <Box>
              <CardActions>
                <Grid container>
                  <Grid item xs={4}>
                    <FormControlLabel
                      control={<Checkbox 
                                  checked={isFinished}
                                  value={isFinished}
                                  onChange={onChangeIsFinished}
                                />} 
                      label="finish"
                      sx={{
                        fontSize: '8px',
                        marginBottom: '3px'
                      }}
                    />
                  </Grid>
                  <Grid item xs={8}>
                    <Button
                      onClick={onClickDelete}
                      sx={{
                        color: 'black',
                        fontSize: '8px',
                        marginBottom: '3px'
                      }}
                    >
                      <DeleteIcon />
                      Delete
                    </Button>
                    <Button
                      onClick={onclickToShowTodo}
                      sx={{
                        color: 'black',
                        fontSize: '8px'
                      }}
                    >
                      <StickyNote2Icon />
                      more
                    </Button>
                  </Grid>
                </Grid>
              </CardActions>
            </Box>
            <Box
              marginY="auto"
            >
              <p>
                {handleToDateFormat(taskCard.createdAt)}
              </p>
            </Box>
          </Box>
        </Box>
      </Box>
      <ShowTaskCardModal 
        open={open}
        onClose={onClickModalClose}
        taskCard={taskCard}
      />
    </>
  )
}

export default IndexTaskCard;