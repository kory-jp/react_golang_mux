import { Button, CardActions, Checkbox, FormControlLabel, Grid, Modal, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { useCallback, useEffect, useState, VFC } from "react";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { TaskCard } from "../../../reducks/taskCards/types";
import EditTaskCardForm from "./EditTaskCardForm";
import { deleteTaskCard, updateIsFinished } from "../../../reducks/taskCards/operations";
import { useDispatch } from "react-redux";

type Props = {
  open: boolean
  onClose: () => void
  taskCard: TaskCard
}

export const ShowTaskCardModal: VFC<Props> = (props) => {
  const dispatch = useDispatch()
  const {open, onClose, taskCard} = props;
  const [isFinished, setIsFinished] = useState(false)
  const [editModalOpen, setEditModalOpen] = useState(false)

  useEffect(() => {
    setIsFinished(taskCard.isFinished)
  }, [])

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
    onClose()
  }, [])


  const onClickOpenEditModal = useCallback(() => {
    setEditModalOpen(true)
  }, [])

  const onClickOpenEditModalClose = useCallback(() => {
    setEditModalOpen(false)
  }, [])

  return(
    <>
      <Modal
        id="showTaskCardModal"
        open={open}
        onClose={onClose}
      >
        <Box
          id="showTCModalContainer"
          sx={{
            backgroundColor: isFinished ? 'gray' : 'white',
            width: {
              xs: "80%",
              md: "60%",
              lg: "50%",
            },
            marginX: "auto",
            marginY: "40px",
            height: "90vh",
            borderRadius: "10px",
            overflow: "hidden"
          }}
        >
          <Box 
            id="showTCModalWrapper"
            margin="40px"
          >
            <Box 
              id="showTcTitle"
              marginBottom="40px"
            >
              <Box>
                <Typography
                  fontSize="32px !important"
                  fontWeight="bold"
                >
                  {taskCard.title}
                </Typography>
              </Box>
              <Box
                textAlign="right"
              >
                <p>タスクカード</p>
              </Box>
            </Box>
            <Box 
              id="showTcPurpose"
              marginBottom="40px"  
            >
              <Box
                marginBottom="16px"
              >
                <Typography
                  fontSize="24px !important"
                  fontWeight="bold !important"
                >
                  目的: なぜこのタスクをする必要があるのか？
                </Typography>
              </Box>
              <Box>
                {taskCard.purpose}
              </Box>
            </Box>
            <Box
              id="showTcContent"
              marginBottom="40px"
            >
              <Box
                marginBottom="16px"
              >
                <Typography
                  fontSize="24px !important"
                  fontWeight="bold !important"
                >
                  作業内容: 具体的にどのような作業をするのか
                </Typography>
              </Box>
              <Box>
                {taskCard.content}
              </Box>
            </Box>
            <Box 
              id="showTcMemo"
              marginBottom="40px"  
            >
              <Box
                marginBottom="16px"
              >
                <Typography
                  fontSize="24px !important"
                  fontWeight="bold !important"
                >
                  メモ
                </Typography>
              </Box>
              <Box>
              {taskCard.memo}
              </Box>            
            </Box>
            <Box id="showTcModalEditSec">
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
                        onClick={onClickOpenEditModal}
                        sx={{
                          color: 'black',
                          fontSize: '8px'
                        }}
                      >
                        <EditIcon />
                        Edit
                      </Button>
                    </Grid>
                  </Grid>
                </CardActions>
            </Box>
          </Box>
        </Box>
      </Modal>
      <EditTaskCardForm 
        open={editModalOpen}
        onClose={onClickOpenEditModalClose}
        taskCard={taskCard}
      />
    </>
  )
}


export default ShowTaskCardModal;