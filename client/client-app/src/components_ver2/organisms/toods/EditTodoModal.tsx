import axios from "axios";
import { ChangeEvent, FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import useReturnTop from "../../../hooks/useReturnTop";
import { nowLoadingState } from "../../../reducks/loading/actions";
import { RootState } from "../../../reducks/store/store";
import { indexTags } from "../../../reducks/tags/operations";
import { Tags } from "../../../reducks/tags/types";
import { pushToast } from "../../../reducks/toasts/actions";
import { updateTodo } from "../../../reducks/todos/operations";
import { Todo } from "../../../reducks/todos/types";
import { makeOptions } from "../../../utils/makeOptions";
import DefaultInputModal from "./DefaultInputModal";

type Params = {
  id: string | undefined
}

type Props = {
  open: boolean,
  onClose: () => void,
}

export const EditTodoModal: FC<Props> = (props) => {
  const { open, onClose} = props
  const dispatch = useDispatch()
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [tags, setTags] = useState<Tags>([])
  const [importance, setImportance] = useState(0)
  const [urgency, setUrgency] = useState(0)
  const [imagePath, setImagePath] = useState('')
  const [image, setImage] = useState<File>()
  const [preview, setPreview] =useState('')
  const params: Params = useParams();
  const id: number = Number(params.id)
  const { importanceOptions, urgencyOptions } = makeOptions()
  const returnTop = useReturnTop()

  const getTodoInfo = useCallback((id: number) => {
      dispatch(nowLoadingState(true))
      const apiURL = process.env.REACT_APP_API_URL + `todos/${id}`
      axios
        .get(apiURL,
        {
          withCredentials: true,
          headers:{
            'Accept': 'application/json',  
            'Content-Type': 'multipart/form-data'
          }
        }
        ).then((response) => {
          if (response.data.status === 200) {
            const todo = response.data.todo
            setTitle(todo.title)
            setContent(todo.content)
            setImportance(todo.importance)
            setUrgency(todo.urgency)
            setTags(todo.tags)
            setImagePath(todo.imagePath)
            let imagePath: string = ""
            if (process.env.NODE_ENV === "production") {
              imagePath = todo.imagePath ? todo.imagePath : ""
            } else {
              imagePath = todo.imagePath ? process.env.REACT_APP_API_URL + `img/${todo.imagePath}` : " "
            }
            setPreview(imagePath)
          } else {
            dispatch(pushToast({title: response.data.message, severity: "error"}))               
          }
        })
        .catch((error) => {
          dispatch(pushToast({title: 'データ取得に失敗しました', severity: "error"}))
        })
        .finally(() => {
          setTimeout(() => {
            dispatch(nowLoadingState(false));
          }, 800);
        });
  }, [dispatch])

  useEffect(() => {
    getTodoInfo(id)
    dispatch(indexTags())
  }, [dispatch, getTodoInfo, id])

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
  
  const previewImage = useCallback((event: ChangeEvent<HTMLInputElement> ) => {
    if (event.target.files === null) {
      return;
    }
    const imageFile = event.target.files[0];
    setPreview(window.URL.createObjectURL(imageFile))
  },[])
  
  const onClickInputImage = useCallback((event: ChangeEvent<HTMLInputElement>  )=> {
    if (event.target.files === null) {
      return;
    }
    const file = event.target.files[0]
    setImage(file)
    previewImage(event)
  }, [setImage, previewImage])

  const onClickCancelImage = useCallback(() => {
    setImage(undefined)
    setImagePath('')
    setPreview('')
  }, [])

  const createFormData = useCallback(() => {
    const formData = new FormData()
    formData.append('id', String(id))
    formData.append('title', title)
    formData.append('content', content)
    formData.append('importance', String(importance))
    formData.append('urgency', String(urgency))
    if (image) formData.append('image', image)
    formData.append('imagePath', imagePath)
    for(let i in tags) {
      let tagId = String(tags[i].id)
      formData.append('tagIds', tagId)
    }
    return formData
  }, [id, title, content, importance, urgency, tags, image, imagePath])


  const formData = createFormData()
  const onClickEditTodo = useCallback(() => {
    dispatch(updateTodo(id, formData))
    returnTop()
    onClose()
  }, [dispatch, returnTop, onClose, id, formData])


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
        onClickSubmitTodo={onClickEditTodo}
        topLabel='タスク更新'
        buttonLabel="更新"
      />      
    </>
  )
}

export default EditTodoModal;