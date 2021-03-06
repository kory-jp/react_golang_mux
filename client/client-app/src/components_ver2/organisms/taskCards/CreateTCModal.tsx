import { FC, useCallback, useState } from "react";
import { createTaskCard } from "../../../reducks/taskCards/operations";
import { useDispatch } from "react-redux";
import { useParams } from "react-router-dom";
import DefaultInputTCModal from "./DefaultInputTCModal";
import usePagination from "../../../hooks/usePagination";
import useReturnTop from "../../../hooks/useReturnTop";
import { useIncompleteTaskCardCount } from "../../../hooks/useIncompleteTaskCardCount";

type Props = {
  open: boolean,
  onClose: () => void,
}

type Params = {
  id: string | undefined
}

export const CreateTCModal: FC<Props> = (props) => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const {open, onClose} = props
  const [title, setTitle] = useState('')
  const [purpose, setPurpose] = useState("")
  const [content, setContent] = useState("")
  const [memo, setMemo] = useState("")
  const {setSumPage, queryPage} = usePagination()
  const returnTop = useReturnTop()
  const {incompleteTaskCardCount, setIncompleteTaskCardCount} = useIncompleteTaskCardCount()

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

  const onClickNewTaskCard = useCallback(() => {
    dispatch(createTaskCard(id, title, purpose, content, memo, setSumPage, queryPage))
    setTitle('')
    setPurpose('')
    setContent('')
    setMemo('')
    onClose()
    returnTop()
    setIncompleteTaskCardCount(incompleteTaskCardCount + 1)
  }, [dispatch, onClose, returnTop, setIncompleteTaskCardCount, setSumPage, incompleteTaskCardCount, queryPage,id, title, purpose, content, memo])

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
        onClickSubmitTC={onClickNewTaskCard}
        topLabel='????????????????????????'
        buttonLabel='??????'
      />
    </>
  )
}

export default CreateTCModal;