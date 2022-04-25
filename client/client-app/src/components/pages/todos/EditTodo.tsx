import axios from "axios";
import ClearIcon from '@mui/icons-material/Clear';
import CameraAltIcon from '@mui/icons-material/CameraAlt';
import { Button, Divider, FormControl, Input, InputLabel, Paper, Stack, TextField} from "@mui/material";
import { Box } from "@mui/system";
import React, { ChangeEvent, FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { nowLoadingState } from "../../../reducks/loading/actions";
import { pushToast } from "../../../reducks/toasts/actions";
import { updateTodo } from "../../../reducks/todos/operations";
import TagSelection from "../../molecules/tag/TagSelection";
import { indexTags } from "../../../reducks/tags/operations";
import { RootState } from "../../../reducks/store/store";
import { Tags } from "../../../reducks/tags/types";
import { makeOptions } from "../../../utils/makeOptions";
import { Todo } from "../../../reducks/todos/types"


type Params = {
  id: string | undefined
}

export const EditTodo: FC = () => {
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
          if (response.data.status == 200) {
            const todo = response.data.todo
            setTitle(todo.title)
            setContent(todo.content)
            setImportance(todo.importance)
            setUrgency(todo.urgency)
            setTags(todo.tags)
            setImagePath(todo.imagePath)
            const imagePath = todo.imagePath? process.env.REACT_APP_API_URL + `img/${todo.imagePath}` : ''            
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
  }, [])

  useEffect(() => {
    getTodoInfo(id)
    dispatch(indexTags())
  }, [getTodoInfo])

  const options = useSelector((state: RootState) => state.tags)

  const inputTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])
  
  const inputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
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
  
  const inputImage = useCallback((event: ChangeEvent<HTMLInputElement>  )=> {
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
  }, [title, content, importance, urgency, tags, image, imagePath])


  const formData = createFormData()
  const onClickEditTodo = useCallback(() => {
    dispatch(updateTodo(id, formData))
  }, [formData])

  return(
    <>
      <Paper
        sx={{
          padding: {
            xs: '10px',
            md: '30px'
          },
          marginX: {
            xs: '0px',
            sm: '80px',
            md: '150px'
          }
        }}
      >
        <Box
          fontWeight='bold'
          marginBottom='10px'
          sx={{
            fontSize: {
              sx: '20px',
              md: '35px'
            }
          }}
        >
          TODO編集
        </Box>
        <Divider />
        <Box
          sx={{
            margin: {
              sx: '10px',
              md: '20px'
            }
          }}
        >
          <Stack spacing={3}>
            <FormControl fullWidth>
              <TextField 
                label='Ttile'
                type='text'
                value={title}
                variant="standard"
                onChange={inputTitle}
              />
            </FormControl>
            <FormControl fullWidth>
              <TextField 
                label='Memo'
                multiline
                type='text'
                value={content}
                variant="standard"
                rows={4}
                onChange={inputContent}
              />
            </FormControl>
            {/* --- タグ選択 --- */}
            <FormControl fullWidth>
              <TagSelection
                placeholder={"タグを選択してください"}  
                isMulti={true} 
                options={options} 
                onChange={onChangeSelectTags}
                values={tags}
              />
            </FormControl>
            <FormControl fullWidth>
              <TagSelection
                placeholder={"重要度を選択してください"} 
                isMulti={false} 
                options={importanceOptions} 
                onChange={onChangeImportance}
                value={importance === 1 ? importanceOptions[0] : importanceOptions[1]}
              />
            </FormControl>
            <FormControl fullWidth>
              <TagSelection
                placeholder={"緊急度を選択してください"} 
                isMulti={false} 
                options={urgencyOptions} 
                onChange={onChangeUrgency}
                value={urgency === 1 ? urgencyOptions[0] : urgencyOptions[1]}
              />
            </FormControl>

            <InputLabel htmlFor="upImage">
              <Input 
                id='upImage'
                name='image'
                type='file'
                inputProps={{accept: "image/*, .jpg, .jpeg, .png"}}
                onChange={inputImage}
                sx={{
                  display: 'none'
                }}
              />
              <Button variant="contained" component='span'>
                <CameraAltIcon />
              </Button>
            </InputLabel>
            {
              preview ?
              <Box>
                <Button>
                  <ClearIcon onClick={onClickCancelImage}/>
                </Button>
                <img 
                  src={preview}
                  alt='preview_img'
                  width='100%'
                  height='auto'
                />
              </Box> : null
            }
            <PrimaryButton
              disabled={title === ''}
              onClick={onClickEditTodo}
            >
              更新
            </PrimaryButton>
          </Stack>
        </Box>
      </Paper>
    </>
  )
}

export default EditTodo;