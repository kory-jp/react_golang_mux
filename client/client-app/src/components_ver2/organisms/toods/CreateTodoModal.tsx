import { ChangeEvent, FC, useCallback, useEffect, useState } from "react";
import { Tags } from "../../../reducks/tags/types";
import { indexTags } from "../../../reducks/tags/operations";
import { RootState } from "../../../reducks/store/store";
import { useDispatch, useSelector } from "react-redux";
import { makeOptions } from "../../../utils/makeOptions";
import { Todo } from "../../../reducks/todos/types";
import { createTodo } from "../../../reducks/todos/operations";
import DefaultInputModal from "./DefaultInputModal";

type Props = {
  open: boolean,
  onClose: () => void,
}

export const CreateTodoModal: FC<Props> = (props) => {
  const dispatch = useDispatch()
  const {open, onClose} = props;
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [tags, setTags] = useState<Tags>([])
  const [importance, setImportance] = useState<number | undefined>(undefined)
  const [urgency, setUrgency] = useState<number | undefined>(undefined)
  const { importanceOptions, urgencyOptions } = makeOptions()
  const [image, setImage] = useState<File>()
  const [preview, setPreview] =useState('')
  
  useEffect(() => {
    dispatch(indexTags())
  }, [dispatch])

  const options = useSelector((state: RootState) => state.tags)

  const onChangeInputTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])

  const onChangeInputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])

  const onChangeSelectTags = useCallback((event: React.SetStateAction<Tags>) => {
    setTags(event)
  }, [setTags])

  const onChangeImportance = useCallback((event: Todo) => {
    setImportance(event.id)
  }, [setImportance])

  const onChangeUrgency = useCallback((event: Todo) => {
    setUrgency(event.id)
  }, [setUrgency])

  const previewImage = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.files === null) {
      return;
    }
    const imageFile = event.target.files[0];
    setPreview(window.URL.createObjectURL(imageFile))
  },[])

  const onClickInputImage = useCallback((event: ChangeEvent<HTMLInputElement>)=> {
    console.log('img OK?')
    if (event.target.files === null) {
      return;
    }
    const file = event.target.files[0]
    setImage(file)
    previewImage(event)
  }, [setImage, previewImage])

  const onClickCancelImage = useCallback(() => {
    setImage(undefined)
    setPreview('')
  }, [])

  const createFormData = useCallback(() => {
    const formData = new FormData()
    formData.append('title', title)
    formData.append('content', content)
    formData.append('importance', String(importance))
    formData.append('urgency', String(urgency))
    if (image) formData.append('image', image)
    for(let i in tags) {
      let tagId = String(tags[i].id)
      formData.append('tagIds', tagId)
    }
    return formData
  }, [title, content, importance, urgency, image, tags])

  const formData = createFormData()

  const onClickNewTodo = useCallback(() => {
    dispatch(createTodo(formData))
    onClose()
  }, [formData])


  return(
    <>
      <DefaultInputModal 
        open={open}
        onClose={onClose}
        title={title}
        content={content}
        tags={tags}
        importance={importance}
        urgency={urgency}
        preview={preview}
        options={options}
        importanceOptions={importanceOptions}
        urgencyOptions={urgencyOptions}
        onChangeInputTitle={onChangeInputTitle}
        onChangeInputContent={onChangeInputContent}
        onChangeSelectTags={onChangeSelectTags}
        onChangeImportance={onChangeImportance}
        onChangeUrgency={onChangeUrgency}
        onClickInputImage={onClickInputImage}
        onClickCancelImage={onClickCancelImage}
        onClickNewTodo={onClickNewTodo}
      />
    </>
  )
}

export default CreateTodoModal;