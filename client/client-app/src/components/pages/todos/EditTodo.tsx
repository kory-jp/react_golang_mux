import axios from "axios";
import ClearIcon from '@mui/icons-material/Clear';
import CameraAltIcon from '@mui/icons-material/CameraAlt';
import { Button, Divider, FormControl, Input, InputLabel, Paper, Stack, TextField} from "@mui/material";
import { Box } from "@mui/system";
import React, { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { nowLoadingState } from "../../../reducks/loading/actions";
import { pushToast } from "../../../reducks/toasts/actions";
import { updateTodo } from "../../../reducks/todos/operations";

type Params = {
  id: string | undefined
}

export const EditTodo: FC = () => {
  const dispatch = useDispatch()
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [imagePath, setImagePath] = useState('')
  const [image, setImage] = useState<File>()
  const [preview, setPreview] =useState('')
  const params: Params = useParams();
  const id: number = Number(params.id)

  const getTodoInfo = useCallback((id: number) => {
      dispatch(nowLoadingState(true))
      axios
        .get(`http://localhost:8000/api/todos/${id}`,
        {
          withCredentials: true,
          headers:{
            'Accept': 'application/json',  
            'Content-Type': 'multipart/form-data'
          }
        }
        ).then((response) => {
          const todo = response.data
          setTitle(todo.title)
          setContent(todo.content)
          setImagePath(todo.imagePath)
          const imagePath = todo.imagePath? `http://localhost:8000/api/img/${todo.imagePath}` : ''
          setPreview(imagePath)
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
  }, [getTodoInfo])

  const inputTitle = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  },[setTitle])
  
  const inputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])
  
  const previewImage = useCallback((event) => {
    const imageFile = event.target.files[0];
    setPreview(window.URL.createObjectURL(imageFile))
  },[])
  
  const inputImage = useCallback((event)=> {
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
    if (image) formData.append('image', image)
    formData.append('imagePath', imagePath)
    return formData
  }, [title, content, image, imagePath])


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