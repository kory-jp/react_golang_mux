import { FormControl, Input, Modal, TextField } from "@mui/material";
import { Box } from "@mui/system";
import { useCallback, useEffect, useState, VFC } from "react";
import { useDispatch } from "react-redux";
import usePagination from "../../../hooks/usePagination";
import {updateTaskCard } from "../../../reducks/taskCards/operations";
import { TaskCard } from "../../../reducks/taskCards/types";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

type Props = {
  open: boolean,
  onClose: () => void,
  taskCard: TaskCard,
}

export const EditTaskCardForm: VFC<Props> = (props) => {
  const {open, onClose, taskCard} = props
  const dispatch = useDispatch()
  const [title, setTitle] = useState("")
  const [purpose, setPurpose] = useState("")
  const [content, setContent] = useState("")
  const [memo, setMemo] = useState("")
  const {setSumPage, queryPage} = usePagination()

  useEffect(() => {
    setTitle(taskCard.title)
    setPurpose(taskCard.purpose)
    setContent(taskCard.content)
    setMemo(taskCard.memo)
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
    dispatch(updateTaskCard(taskCard.id, taskCard.todoId, title, purpose, content, memo, setSumPage, queryPage))
    onClose()
  }, [taskCard, title, purpose, content, memo])

  return(
    <Modal
      id="editTcModal"
      open={open}
      onClose={onClose}
    >
      <Box
        id="editTcModalContainer"
        sx={{
          backgroundColor: "white",
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
          id="modalWrapper"
          margin="40px"
        >
          <Box
            sx={{
              display: "flex",
              marginBottom: "80px"
            }}
          >
            <Box
              marginRight="16px"
            >
              <Input 
                placeholder="タイトル"
                sx={{
                  width: "240px"
                }}
                value={title}
                onChange={onChangeTitle}
              />
            </Box>
            <Box>
              タスクカード
            </Box>
          </Box>
          <Box
            marginBottom="40px"
          >
            <FormControl id="editTcPurposeFormControl" fullWidth>
              <TextField
              label="目的: なぜこのタスクをする必要があるのか？"
              multiline
              rows={5}
              variant="standard"
              value={purpose}
              onChange={onChangePurpose}
              />
            </FormControl>
          </Box>
          <Box
            marginBottom="40px"
          >
            <FormControl id="editTcContentFormControl" fullWidth>
              <TextField
              label="作業内容: 具体的にどのような作業をするのか"
              multiline
              rows={5}
              variant="standard"
              value={content}
              onChange={onChangeContent}
              />
            </FormControl>
          </Box>
          <Box
            marginBottom="80px"
          >
            <FormControl id="formControl" fullWidth>
              <TextField
              label="メモ"
              multiline
              rows={5}
              variant="standard"
              value={memo}
              onChange={onChangeMemo}
              />
            </FormControl>
          </Box>
          <Box
            display="flex"
            justifyContent="center"
          >
            <Box
              marginRight="24px"
            >
              <PrimaryButton
                onClick={onClickNewTaskCard}
              >
                タスク編集
              </PrimaryButton>
            </Box>
            <Box>
              <PrimaryButton
                onClick={onClose}
              >
                戻る
              </PrimaryButton>
            </Box>
          </Box>
        </Box>
      </Box>
    </Modal>
  )
}

export default EditTaskCardForm;