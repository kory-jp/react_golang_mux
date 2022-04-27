import { Box, Button, CardActions, Checkbox, FormControlLabel, Grid } from "@mui/material";
import { useCallback, useEffect, useState, VFC } from "react";
import DeleteIcon from '@mui/icons-material/Delete';
import StickyNote2Icon from '@mui/icons-material/StickyNote2';
import { TaskCard } from "../../../reducks/taskCards/types";
import ShowTaskCardModal from "./ShowTaskCardModal";

type Props ={
  taskCard: TaskCard
}

export const IndexTaskCard: VFC<Props> = (props) => {
  const {taskCard} = props
  const [isFinished, setIsFinished] = useState(false)
  const [finish, setFinish] = useState(false)
  const [open, setOpen] = useState(false);
  useEffect(() => {
    setIsFinished(taskCard.isFinished)
  },[])

  const onChangeIsFinished = useCallback(() => {
    if (isFinished) {
      setIsFinished(false)
      // dispatch(updateIsFinished(id, false))
    } else {
      setIsFinished(true)
      // dispatch(updateIsFinished(id, true))
    }
  }, [isFinished])

  const onClickDelete = useCallback(() => {
    console.log("delete!")
  }, [])

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
                                  checked={finish}
                                  value={finish}
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
                {taskCard.createdAt}
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