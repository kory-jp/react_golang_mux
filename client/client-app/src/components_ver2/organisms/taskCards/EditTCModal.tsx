import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import usePagination from "../../../hooks/usePagination";
import { updateTaskCard } from "../../../reducks/taskCards/operations";
import { TaskCard } from "../../../reducks/taskCards/types";
import DefaultInputTCModal from "./DefaultInputTCModal";

type Props = {
  open: boolean,
  onClose: () => void,
  taskCard: TaskCard,
}

export const EditTCModal: FC<Props> = (props) => {
  const {open, onClose, taskCard} = props
  const dispatch = useDispatch()
  const [title, setTitle] = useState('')
  const [purpose, setPurpose] = useState("")
  const [content, setContent] = useState("")
  const [memo, setMemo] = useState("")
  const {setSumPage, queryPage} = usePagination()

  useEffect(() => {
    setTitle(taskCard.title)
    setPurpose(taskCard.purpose)
    setContent(taskCard.content)
    setMemo(taskCard.memo)
  }, [taskCard.content, taskCard.memo, taskCard.purpose, taskCard.title])

  const onChangeInputTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])

  const onChangeInputPurpose = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPurpose(event.target.value)
  },[setPurpose])

  const onChangeInputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])

  const onChangeInputMemo = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setMemo(event.target.value)
  },[setMemo])

  const onClickUpdateTaskCard = useCallback(() => {
    dispatch(updateTaskCard(taskCard.id, taskCard.todoId, title, purpose, content, memo, setSumPage, queryPage))
    onClose()
  }, [dispatch, onClose, setSumPage, queryPage, taskCard, title, purpose, content, memo])

  return(
    <>
      <DefaultInputTCModal 
        open={open} 
        onClose={onClose}
        title={title}
        purpose={purpose}
        content={content}
        memo={memo}
        onChangeInputTitle={onChangeInputTitle}
        onChangeInputPurpose={onChangeInputPurpose}
        onChangeInputContent={onChangeInputContent}
        onChangeInputMemo={onChangeInputMemo}
        onClickSubmitTC={onClickUpdateTaskCard}
        topLabel='タスクカード更新'
        buttonLabel='更新'
      />    
    </>
  )
}

export default EditTCModal;