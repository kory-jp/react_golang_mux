import { Box, FormControl, Input, Modal, TextField } from "@mui/material";
import { VFC } from "react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

type Props = {
  open: boolean
  onClose: () => void
  title: string
  purpose: string
  content: string
  memo: string
  onChangeTitle: (event: React.ChangeEvent<HTMLInputElement>) => void
  onChangePurpose: (event: React.ChangeEvent<HTMLInputElement>) => void
  onChangeContent: (event: React.ChangeEvent<HTMLInputElement>) => void
  onChangeMemo: (event: React.ChangeEvent<HTMLInputElement>) => void
  onClickNewTaskCard: () => void
}

export const InputTaskCardForm: VFC<Props> = (props) => {
  const {
    open, 
    onClose, 
    title, 
    purpose, 
    content, 
    memo, 
    onChangeTitle, 
    onChangePurpose, 
    onChangeContent, 
    onChangeMemo, 
    onClickNewTaskCard
  } = props

  return (
    <Modal
      id="modal"
      open={open}
      onClose={onClose}
    >
      <Box
        id="modalContainer"
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
            <FormControl id="formControl" fullWidth>
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
            <FormControl id="formControl" fullWidth>
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
            textAlign="center"
          >
            <PrimaryButton
              onClick={onClickNewTaskCard}
            >
              タスク追加
            </PrimaryButton>
          </Box>
        </Box>
      </Box>
    </Modal>
  )
}

export default InputTaskCardForm